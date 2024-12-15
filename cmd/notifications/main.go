package main

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/notifications"
	notificationsServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/notifications"
	"log"

	notificationsApp "github.com/go-park-mail-ru/2024_2_kotyari/internal/apps/notifications"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	//notificationsDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/notifications"
	//errResolveLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	//"github.com/gorilla/mux"
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
	slogLogger := logger.InitLogger()
	//errResolver := errResolveLib.NewErrorStore()
	db, err := postgres.LoadPgxPool()
	if err != nil {
		log.Fatal(err)
	}

	notificationsRepo := notifications.NewNotificationsStore(db, slogLogger)
	notificationsWorker := notificationsServiceLib.NewNotificationsWorker(notificationsRepo, slogLogger)
	//notificationDelivery := notificationsDeliveryLib.NewNotificationsDelivery(errResolver, slogLogger)
	notificationApp := notificationsApp.NewNotificationsApp(config, notificationsWorker, slogLogger)

	notificationApp.Run()
}
