package config

import (
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"
)

type EnvConfig struct {
	SQL_DB_HOST string `env:"SQL_DB_HOST" envDefault:"localhost"`
	SQL_DB_PORT int32  `env:"SQL_DB_PORT" envDefault:"5432"`
	SQL_DB_USER string `env:"SQL_DB_USER" envDefault:"postgres"`
	SQL_DB_PASS string `env:"SQL_DB_PASS" envDefault:"password"`

	SQLITE_PATH string `env:"SQLITE_PATH" envDefault:"sqlite.db"`

	PORT      int32  `env:"PORT" envDefault:"8000"`
	LOG_LEVEL string `env:"LOG_LEVEL" envDefault:"info"`

	COGNITO_POOL_ID   string `env:"COGNITO_POOL_ID" envDefault:""`
	COGNITO_CLIENT_ID string `env:"COGNITO_CLIENT_ID" envDefault:""`
}

var envConfig EnvConfig
var envOnce sync.Once

func ParseEnvironment() *EnvConfig {
	envOnce.Do(func() {
		err := env.Parse(&envConfig)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
	})

	return &envConfig
}
