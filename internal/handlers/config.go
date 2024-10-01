package handlers

import (
	"log"

	"github.com/caarlos0/env"
)

type session struct {
	SessionKey  string `env:"SESSION_KEY"`  // Ключ сессии из переменной окружения
	SessionName string `env:"SESSION_NAME"` // Имя сессии из переменной окружения
}

func initSessions() session {
	var config session

	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("error parsing environment variables: %v", err)
	}

	return config
}

func initTestSession() session {
	return session{
		SessionKey:  "test-key",
		SessionName: "test-name",
	}
}
