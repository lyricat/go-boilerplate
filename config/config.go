package config

import (
	"github.com/fox-one/pkg/store/db"
)

type (
	Config struct {
		DB db.Config `yaml:"db"`
	}
)
