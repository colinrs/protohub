package config

import (
	"github.com/colinrs/protohub/internal/infra"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DataBase     *infra.DBConfig `json:"Database" yaml:"Database"`
	PasswdSecret string          `json:"PasswdSecret" yaml:"PasswdSecret"`
	JwtSecret    string          `json:"JwtSecret" yaml:"JwtSecret"`
	JwtExpired   int64           `json:"JwtExpired" yaml:"JwtExpired"`
	Redis        redis.RedisConf `json:"Redis" yaml:"Redis"`
}
