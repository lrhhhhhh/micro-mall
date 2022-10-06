## intro
这是一个用 go-zero 实现的商城秒杀项目。
- 使用 redis + lua 抗住大部分的并发请求
- 库存扣减和创建订单: 两种实现，分别使用分布式事务库dtm的 TCC 模式、二阶段消息实现，保证数据的最终一致性
- 使用 kafka 实现的误差1秒以内的延迟队列实现15分钟未支付关闭订单
- 服务发现：本地开发使用etcd，生产环境中直接使用kubernetes
- trace： jaeger
- metric：prometheus + grafana  
- api gateway：主要用来做统一的鉴权中心，使用go-zero的api文件实现， 使用jwt的方式做鉴权 

## 项目中如何使用dtm的tcc模式和二阶段消息模式
TCC模式下

二阶段消息下：

## 性能 


## 运行 
- 详细内容请参考 Makefile 中的注释
- 请使用提供的 `postman_collection.json` 进行创建新数据并测试

### 使用 docker-compose 进行本地开发
使用 docker-compose up -d 运行mysql、redis、etcd、jaeger、kafka、dtm 等依赖软件  
编写的程序则在本地运行，省去构建镜像时间。详细步骤如下：   
```shell
cd micro-mall
make up                     # docker-compose up -d 
make docker-compose-initdb  # 创建数据库、表
make run                    # 运行所有程序
make stop                   # 关闭所有程序 
make down                   # docker-compose down 
```


### 在 minikube 中运行
```shell
make bulid-images
make load 
make minikube-ingress 
make apply 
```

在本地构建镜像后，直接使用load命令加载到minikube中，省去 push 到 registry 又 pull 下来的消耗。

**注意**，部分服务使用init-container进行了启动顺序控制，首次创建时，数据库因为没有初始化会使得依赖数据库的服务启动失败。  
这需要在 `make apply` 之后使用 `make initdb` 命令初始化数据库，运行命令后，按回车进入mysql交互命令行，将dummy.sql中的内容粘贴并运行即可


## 测试 
### 简单测试 api 
使用 postman 打开项目里的postman的json文件

### 压力测试 
使用wrk进行压力测试，命令为`make wrk_test`，详细参考 `wrk.lua` 文件

## TODO 
