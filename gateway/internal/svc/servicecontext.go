package svc

import (
	"gateway/internal/config"
	"gateway/internal/middleware"
	"gateway/service/activity/activityclient"
	"gateway/service/seckill/seckillclient"
	"gateway/service/stock/stockclient"
	"gateway/service/user/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	UserRpc        userclient.User
	SeckillRpc     seckillclient.Seckill
	StockRpc       stockclient.Stock
	ActivityRpc    activityclient.Activity
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		SeckillRpc:     seckillclient.NewSeckill(zrpc.MustNewClient(c.SeckillRpcConf)),
		StockRpc:       stockclient.NewStock(zrpc.MustNewClient(c.StockRpcConf)),
		ActivityRpc:    activityclient.NewActivity(zrpc.MustNewClient(c.ActivityRpcConf)),
	}
}
