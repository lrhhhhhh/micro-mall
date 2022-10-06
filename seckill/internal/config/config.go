package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Redis redis.RedisConf
	DB    struct {
		DataSource string
	}
	DtmServer struct {
		Addr string
	}

	StockRpcConf zrpc.RpcClientConf
	OrderRpcConf zrpc.RpcClientConf

	StockDeductLua      string
	StockDeductSha1     string
	StockRedisKeyFormat string
}
