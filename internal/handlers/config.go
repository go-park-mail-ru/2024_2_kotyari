package handlers

import (
	"log"

	"github.com/caarlos0/env"
)

type sessionConfig struct {
	SessionKey  string `env:"SESSION_KEY"`  // Ключ сессии из переменной окружения
	SessionName string `env:"SESSION_NAME"` // Имя сессии из переменной окружения
}

func initSessions() sessionConfig {
	var config sessionConfig

	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}

	return config
}

func initTestSession() sessionConfig {
	return sessionConfig{
		SessionKey:  "test-key",
		SessionName: "test-name",
	}
}
