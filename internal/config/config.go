package config

import (
	"github.com/colinrs/protohub/internal/infra"

	"github.com/zeromicro/go-zero/rest"
)

// Auth JWT 认证需要的密钥和过期时间配置
type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type Config struct {
	rest.RestConf
	Auth     Auth
	DataBase *infra.DBConfig `json:"Database" yaml:"Database"`
}
