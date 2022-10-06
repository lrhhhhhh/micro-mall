package svc

import (
	"database/sql"
	redisv8 "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"seckill/internal/config"
	"seckill/internal/model"
	"seckill/service/order"
	"seckill/service/order/orderclient"
	"seckill/service/stock"
	"seckill/service/stock/stockclient"
)

type ServiceContext struct {
	Config        config.Config
	Redis         *redis.Redis
	RawRedis      *redisv8.Client
	StockRpc      stock.StockClient
	OrderRpc      order.OrderClient
	Mysql         *sql.DB
	ActivityModel model.ActivityModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	db, err := conn.RawDB()
	if err != nil {
		panic(err)
	}

	opts := redisv8.Options{
		Addr:     c.Redis.Host,
		PoolSize: 32,
	}

	return &ServiceContext{
		Config:        c,
		Redis:         redis.New(c.Redis.Host),
		RawRedis:      redisv8.NewClient(&opts),
		StockRpc:      stockclient.NewStock(zrpc.MustNewClient(c.StockRpcConf)),
		OrderRpc:      orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		ActivityModel: model.NewActivityModel(conn),
		Mysql:         db,
	}
}
