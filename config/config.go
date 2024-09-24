package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	SessionKey    string `env:"SESSION_KEY"`    // Ключ сессии из переменной окружения
	SessionName   string `env:"SESSION_NAME"`   // Имя сессии из переменной окружения
	ServerAddress string `env:"SERVER_ADDRESS"` // Адрес сервера из переменной окружения
}

var Cfg Config

func Init() error {
	err := env.Parse(&Cfg)
	if err != nil {
		return fmt.Errorf("error parsing environment variables: %w", err)
	}

	fmt.Printf("ServerAddress: %s\n", Cfg.ServerAddress)
	return nil
}

func GetServerAddress() string {
	return Cfg.ServerAddress
}

func GetSessionName() string {
	return Cfg.SessionName
}
