package config

import (
	"github.com/caarlos0/env"
	"log"
)

type Server struct {
	ServerAddress string `env:"SERVER_ADDRESS"` // Адрес сервера из переменной окружения
}

type Session struct {
	SessionKey  string `env:"SESSION_KEY"`  // Ключ сессии из переменной окружения
	SessionName string `env:"SESSION_NAME"` // Имя сессии из переменной окружения
}

func InitSessions() Session {
	var config Session

	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}

	return config
}

func InitTestSession() Session {
	return Session{
		SessionKey:  "test-key",
		SessionName: "test-name",
	}
}

func InitServerConfig() (Server, error) {
	var config Server

	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}

	return config, nil
}
