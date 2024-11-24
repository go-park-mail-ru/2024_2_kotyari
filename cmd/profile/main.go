package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/apps/profile_service"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/joho/godotenv"
)

const (
	profilesService = "profile_go"
)

func main() {
	err := godotenv.Load(configs.EnvPath)
	if err != nil {
		log.Fatal("[ Error ] отсутствует .env файл")
	}

	viper, err := configs.SetupViper()
	if err != nil {
		log.Fatal("Error setting up viper")
	}

	conf := viper.GetStringMap(profilesService)

	app, err := profile_service.NewProfilesApp(conf)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
