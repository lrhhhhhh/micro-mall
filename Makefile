.PHONY:
goctl:
	@go install github.com/zeromicro/go-zero/tools/goctl@latest   # install goctl
	@goctl env check -i -f --verbose                              # install protoc, proto-gen-go, proto-gen-go-grpc


########################################################################################################################
#################################### command for docker-compose, local dev #############################################
########################################################################################################################

# 通过docker-compose的方式运行依赖软件：redis、mysql、etcd、zookeeper、kafka
up:
	@cd ./deploy/dev && mkdir -p data && cd data && mkdir -p kafka zookeeper redis mysql && chmod 777 kafka zookeeper redis mysql
	@cd ./deploy/dev && pwd && docker-compose up -d


down:
	@cd ./deploy/dev && pwd && docker-compose down


docker-compose-initdb:   # 必须保证mysql容器正在运行，保证mysql容器名字为dev_mysql_1，账号和密码和容器的一致
	@docker cp ./dummy.sql dev_mysql_1:/
	@docker exec dev_mysql_1 sh -c 'mysql -uroot -pSecretKey < dummy.sql'


rm-none:
	@docker rmi $$(docker images | grep "none" | awk '{print $3}')


# 运行程序，需要注意程序可能启动失败，通过ps -ef | grep 'go run' 查看
# 查看gateway程序是否正确运行: lsof -i:8082
# 查看日志文件：nohup.out 和 logs
run:
	@cd ./activity && bash -c "nohup go run activity.go > /dev/null 2>&1 &"
	@cd ./user && bash -c "nohup go run user.go > /dev/null 2>&1 &"
	@cd ./delayqueue && bash -c "ulimit -n 65535 && nohup go run ./cmd/delayqueue.go > nohup.out 2>&1 &"
	@cd ./stock && bash -c "nohup go run stock.go > /dev/null 2>&1 &"
	@cd ./order && bash -c "nohup go run order.go > /dev/null 2>&1 &"
	@sleep 2s
	@cd ./seckill && bash -c "nohup go run seckill.go > /dev/null 2>&1 &"
	@sleep 3s
	@cd ./gateway && bash -c "nohup go run gateway.go > /dev/null 2>&1 &"


# 停止所有程序
stop:
	@lsof -i tcp:8081 | awk '/activity/ {print $$2}' | xargs -r kill
	@lsof -i tcp:8082 | awk '/gateway/ {print $$2}' | xargs  -r kill
	@lsof -i tcp:8083 | awk '/order/ {print $$2}' | xargs -r kill
	@lsof -i tcp:8084 | awk '/seckill/ {print $$2}' | xargs -r kill
	@lsof -i tcp:8085 | awk '/stock/ {print $$2}' | xargs -r kill
	@lsof -i tcp:8086 | awk '/user/ {print $$2}' | xargs -r kill
	@ps -A -ww | awk '/delayqueue/ {print $$1}' | xargs -r kill


# 删除本地开发的日志
delete-logs:
	@cd ./activity && rm -rf ./logs
	@cd ./gateway && rm -rf ./logs
	@cd ./order && rm -rf ./logs
	@cd ./seckill && rm -rf ./logs
	@cd ./stock && rm -rf ./logs
	@cd ./user && rm -rf ./logs
	@cd ./delayqueue && rm -f nohup.out


# 使用wrk进行压力测试
wrk_test:
	@docker run --rm --network="host" -it -v $(PWD):/data skandyla/wrk -d30s -t8 -c400   -s wrk.lua http://localhost:8082/seckill2


########################################################################################################################
############################################ command for minikube ######################################################
########################################################################################################################

# 构建镜像，每个镜像分阶段构建，中间阶段会产生GB级别的none镜像，如果磁盘不够，可以把注释去掉，每构建一个镜像就清理一次。
build-images:
	@#docker rmi $$(docker images | grep "none" | awk '{print $3}') || true
	@cd ./activity && docker build . -t activity:latest
	@#docker rmi $$(docker images | grep "none" | awk '{print $3}') || true
	@cd ./gateway  && docker build . -t gateway:latest
	@#docker rmi $$(docker images | grep "none" | awk '{print $3}') || true
	@cd ./order    && docker build . -t order:latest
	@#docker rmi $$(docker images | grep "none" | awk '{print $3}') || true
	@cd ./seckill  && docker build . -t seckill:latest
	@#docker rmi $$(docker images | grep "none" | awk '{print $3}') || true
	@cd ./stock    && docker build . -t stock:latest
	@#docker rmi $$(docker images | grep "none" | awk '{print $3}') || true
	@cd ./user     && docker build . -t user:latest


# 将需要的镜像下载到宿主机的docker中，然后将所有需要的的镜像load到minikube容器中
load:
	@echo "running..."
	@minikube image rm activity:latest || true
	@minikube image rm gateway:latest  || true
	@minikube image rm order:latest    || true
	@minikube image rm seckill:latest  || true
	@minikube image rm stock:latest    || true
	@minikube image rm user:latest     || true

	@echo "still running..."
	@docker pull mysql:8.0.30
	@docker pull redis:7.0.4
	@docker pull yedf/dtm:1.16
	@docker pull jaegertracing/all-in-one:1.32
	@minikube image load mysql:8.0.30
	@minikube image load redis:7.0.4
	@minikube image load yedf/dtm:1.16
	@minikube image load jaegertracing/all-in-one:1.32

	@echo "still running..."
	@minikube image load activity:latest
	@minikube image load gateway:latest
	@minikube image load order:latest
	@minikube image load seckill:latest
	@minikube image load stock:latest
	@minikube image load user:latest
	@echo "done!"


# 在minikube上运行所有程序
apply:
	@cd ./deploy/prod && minikube kubectl -- apply -f namespace.yaml
	@cd ./deploy/prod && minikube kubectl -- apply -f configMap.yaml
	@cd ./deploy/prod && minikube kubectl -- apply -f serviceAccount.yaml
	@cd ./deploy/prod && minikube kubectl -- apply -f ingress.yaml       # minikube需要先安装ingress

#	@cd ./deploy/prod && minikube kubectl -- apply -f mysqlService.yaml
#	@cd ./deploy/prod && minikube kubectl -- apply -f redisService.yaml
#	@cd ./deploy/prod && minikube kubectl -- apply -f jaegerService.yaml
#	@cd ./deploy/prod && minikube kubectl -- apply -f dtmService.yaml

#	@cd ./deploy/prod && minikube kubectl -- apply -f activityService.yaml
#	@cd ./deploy/prod && minikube kubectl -- apply -f userService.yaml
#	@cd ./deploy/prod && minikube kubectl -- apply -f stockService.yaml
#	@cd ./deploy/prod && minikube kubectl -- apply -f orderService.yaml
#	@cd ./deploy/prod && minikube kubectl -- apply -f seckillService.yaml
#	@cd ./deploy/prod && minikube kubectl -- apply -f gatewayService.yaml


# 给 minikube 开启 ingress
minikube-ingress:
	@minikube addons enable ingress


# install strimzi kafka operator (注意是latest)
kafka-cluster:
	@minikube kubectl -- create namespace kafka --dry-run=client -o yaml | kubectl apply -f -
	@minikube kubectl -- create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
	@cd ./deploy/prod && minikube kubectl -- apply -f kafka-persistent-single.yaml -n kafka
	@minikube kubectl -- wait kafka/kafka-cluster --for=condition=Ready -n kafka


test-kafka-producer:
	@minikube kubectl -- -n kafka run kafka-producer -ti --image=quay.io/strimzi/kafka:0.31.1-kafka-3.2.3 --rm=true --restart=Never -- bin/kafka-console-producer.sh --bootstrap-server kafka-cluster-kafka-bootstrap:9092 --topic my-topic


test-kafka-consumer:
	@minikube kubectl -- -n kafka run kafka-consumer -ti --image=quay.io/strimzi/kafka:0.31.1-kafka-3.2.3 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server kafka-cluster-kafka-bootstrap:9092 --topic my-topic --from-beginning

# 初始化数据库
# 注意，执行命令后，根据提示按下回车键，然后将dummy.sql中的文本内容粘贴到mysql-cli中并执行
initdb:
	@minikube kubectl -- run -it --rm --image=mysql:8.0.30 --restart=Never mysql-client -n micro-mall -- mysql -h mysql-service -pSecretKey


# 删除minikube中命名空间micro-mall内的所有内容
delete-micromall:
	@minikube kubectl -- delete all --all -n=micro-mall

info:
	@minikube kubectl -- get pod -n=micro-mall
	@minikube kubectl -- get svc -n=micro-mall
	@minikube kubectl -- get ingress -n=micro-mall


minikube-delete-pv:     # 删除pv和pvc，仅在测试用使用
	@minikube kubectl -- delete pv --all


minikube-delete-pvc:
	@minikube kubectl -- delete pvc --all
