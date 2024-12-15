package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/apps/notifications"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/joho/godotenv"
)

const configFile = ".env"
const notificationService = "notifications_go"

func main() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	v, err := configs.SetupViper()
	if err != nil {
		log.Fatal("Error loading viper", err.Error())
	}

	config := v.GetStringMap(notificationService)

	notificationsApp := notifications.NewNotificationsApp(config)
	if err = notificationsApp.Run(); err != nil {
		log.Fatal(err)
	}
}
