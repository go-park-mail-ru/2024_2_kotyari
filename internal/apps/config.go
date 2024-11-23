package apps

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
