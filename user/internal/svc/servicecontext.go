package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/config"
	"user/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
