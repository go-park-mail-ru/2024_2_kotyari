package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/apps/rating_updater"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/joho/godotenv"
)

// ratingUpdaterService / TODO: Вынести в другое место
const ratingUpdaterService = "rating_updater"

func main() {
	err := godotenv.Load(configs.EnvPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	v, err := configs.SetupViper()
	if err != nil {
		log.Fatal("Error loading viper", err.Error())
	}

	config := v.GetStringMap(ratingUpdaterService)

	ratingUpdaterApp, err := rating_updater.NewApp(config)
	if err != nil {
		log.Fatal(err)
	}

	err = ratingUpdaterApp.Run()
	if err != nil {
		log.Fatal(err)
	}
}
