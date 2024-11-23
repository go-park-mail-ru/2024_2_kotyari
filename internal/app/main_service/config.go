<<<<<<<< HEAD:internal/apps/config.go
package apps
========
package main_service
>>>>>>>> d5de27b ([HACK-2][improve] микросервис csat):internal/apps/main_service/config.go

import (
	"log"

	"github.com/caarlos0/env"
)

type config struct {
	ServerAddress   string `env:"SERVER_ADDRESS"` // Адрес сервера из переменной окружения
	SessionLifetime string `env:"SESSION_LIFETIME" envDefault:"3600"`
}

func initServer() config {
	var cfg config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}

	return cfg
}
