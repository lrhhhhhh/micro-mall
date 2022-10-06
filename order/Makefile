.PHONY: proto
proto:
	@goctl rpc protoc ./service/order/order.proto --go_out=./service --go-grpc_out=./service --zrpc_out=.

.PHONY: model
model:
	@goctl model mysql datasource --url="root:SecretKey@tcp(127.0.0.1:3306)/order" --table="order" -dir ./internal/model