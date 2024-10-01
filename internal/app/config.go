package app

import (
	"log"

	"github.com/caarlos0/env"
)

type server struct {
	ServerAddress   string `env:"SERVER_ADDRESS"` // Адрес сервера из переменной окружения
	SessionLifetime string `env:"SESSION_LIFETIME" envDefault:"3600"`
}

func initServer() server {
	var config server

	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}

	return config
}
