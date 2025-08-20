package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"
)

type envConfig struct {
	SQL_DB_HOST string `env:"SQL_DB_HOST" envDefault:"localhost"`
	SQL_DB_PORT int32  `env:"SQL_DB_PORT" envDefault:"5432"`
	SQL_DB_USER string `env:"SQL_DB_USER" envDefault:"postgres"`
	SQL_DB_PASS string `env:"SQL_DB_PASS" envDefault:"password"`
	APP_PORT    int32  `env:"APP_PORT" envDefault:"8000"`
}

func ParseEnvironment() *envConfig {
	envConfig := envConfig{}
	err := env.Parse(&envConfig)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	return &envConfig
}
