package svc

import (
	"activity/internal/config"
	"activity/internal/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ActivityModel model.ActivityModel
	Redis         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ActivityModel: model.NewActivityModel(sqlx.NewMysql(c.DB.DataSource)),
		Redis:         redis.New(c.Redis.Host),
	}
}
