package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Redis redis.RedisConf

	ActivityRedisKeyFormat string
	HistoryRedisKeyFormat  string
	StockRedisKeyFormat    string
}
