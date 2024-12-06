package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/apps/promocodes"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/joho/godotenv"
)

const (
	promocodesService = "promocodes_go"
	kafka             = "kafka"
)

func main() {
	err := godotenv.Load(configs.EnvPath)
	if err != nil {
		log.Fatal("[ Error ] отсутствует .env файл")
	}

	v, err := configs.SetupViper()
	if err != nil {
		log.Fatal("Error setting up viper")
	}

	kafkaConf := v.GetStringMap(kafka)
	appConf := v.GetStringMap(promocodesService)

	app, err := promocodes.NewPromoCodesApp(kafkaConf, appConf)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
