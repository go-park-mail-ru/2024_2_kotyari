package main

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/app/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	configs "github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/user"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
)

const (
	userService = "user_go"
	configFile  = ".env"
)

// todo вынос в app
func main() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	viper, err := configs.SetupViper()
	if err != nil {
		log.Fatalf("Error setting up viper %v", err)
	}

	conf := viper.GetStringMap(userService)
	//
	//log.Println(conf)
	server := grpc.NewServer()
	slogLog := logger.InitLogger()

	app, err := user.NewUsersApp(slogLog, server, conf)

	err = app.Run()
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
