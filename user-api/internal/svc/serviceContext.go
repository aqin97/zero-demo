package svc

import (
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserNicknameModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserNicknameModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
