.PHONY: model
model:
	@goctl model mysql datasource --url="root:SecretKey@tcp(127.0.0.1:3306)/user" --table="user" -dir ./internal/model


.PHONY: proto
proto:
	@goctl rpc protoc ./service/user/user.proto --go_out=./service --go-grpc_out=./service --zrpc_out=.