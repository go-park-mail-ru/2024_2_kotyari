package main

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/apps/csat_service"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
)

const csatService = "csat_service"

func main() {
	err := godotenv.Load(configs.EnvPath)
	if err != nil {
		log.Fatal("[ Error ] отсутствует .env файл")
	}

	viper, err := configs.SetupViper()
	if err != nil {
		log.Fatal("Error setting up viper")
	}

	conf := viper.GetStringMap(csatService)

	server := grpc.NewServer()
	logInit := logger.InitLogger()

	app, err := csat_service.NewCsatApp(server, conf, logInit)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
