package main

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/apps/promocodes"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(configs.EnvPath)
	if err != nil {
		log.Fatal("[ Error ] отсутствует .env файл")
	}
	//
	//v, err := configs.SetupViper()
	//if err != nil {
	//	log.Fatal("Error setting up viper")
	//}
	//
	//conf := v.GetStringMap(profilesService)

	app, err := promocodes.NewPromoCodesApp()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
