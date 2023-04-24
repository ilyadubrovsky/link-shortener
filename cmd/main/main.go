package main

import (
	"github.com/joho/godotenv"
	"link-shortener/internal/app"
	"link-shortener/internal/config"
	"link-shortener/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	if err := godotenv.Load(); err != nil {
		logger.Panic(err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Panic(err)
	}

	a, err := app.New(cfg)
	if err != nil {
		logger.Panic(err)
	}

	if err = a.Run(); err != nil {
		logger.Panic(err)
	}
}
