package svc

import (
	"github.com/colinrs/protohub/internal/config"
	"github.com/colinrs/protohub/internal/infra"
	"github.com/colinrs/protohub/internal/middleware"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	DB        *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthMiddleware().Handle,
		DB:        initDB(c),
	}
}

func initDB(c config.Config) *gorm.DB {
	db, err := infra.Database(c.DataBase)
	logx.Must(err)
	return db
}
