package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"link-shortener/pkg/logging"
	"sync"
)

type Config struct {
	StorageConnect string `env:"STORAGE_CONNECT" env-default:""`
	Postgres       Postgres
	ServerPort     string `env:"SERVER_PORT" env-default:"8080"`
	TransportType  string `env:"TRANSPORT_TYPE" env-default:"http"`
}

type Postgres struct {
	Username string `env:"PG_USERNAME" env-default:"postgres"`
	Password string `env:"PG_PASSWORD" env-default:"12345678"`
	Host     string `env:"PG_HOST" env-default:"localhost"`
	Port     string `env:"PG_PORT" env-default:"5432"`
	Database string `env:"PG_DATABASE" env-default:"link-shortener"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() (*Config, error) {
	var err error = nil

	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err = cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Debug(help)
			return
		}
	})

	return instance, err
}
