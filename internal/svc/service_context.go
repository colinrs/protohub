package svc

import (
	"github.com/colinrs/protohub/internal/config"
	"github.com/colinrs/protohub/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware
	DB     *gorm.DB
}

func NewServiceContext(c config.Config, db *gorm.DB) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware().Handle,
		DB:     db,
	}
}
