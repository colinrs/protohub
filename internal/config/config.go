package config

import (
	"github.com/colinrs/protohub/internal/infra"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	DataBase *infra.DBConfig `json:"Database" yaml:"Database"`
}
