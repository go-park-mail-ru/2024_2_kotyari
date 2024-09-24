package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	ServerAddress string // Адрес сервера из переменной окружения
	SessionKey    string // Ключ сессии из переменной окружения
	SessionName   string // Имя сессии из переменной окружения
)

func init() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки файла .env:", err)
	}
	ServerAddress = os.Getenv("SERVER_ADDRESS") // Адрес сервера из переменной окружения
	SessionKey = os.Getenv("SESSION_KEY")       // Ключ сессии из переменной окружения
	SessionName = os.Getenv("SESSION_NAME")
}
