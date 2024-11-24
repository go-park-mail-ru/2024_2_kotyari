package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/app/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	configs "github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
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

	router := mux.NewRouter()
	router.PathPrefix("/metrics").Handler(promhttp.Handler())
	serverProm := http.Server{Handler: router, Addr: fmt.Sprintf(":%d", 8082), ReadHeaderTimeout: 10 * time.Second}

	go func() {
		if err := serverProm.ListenAndServe(); err != nil {
			log.Println("fail auth.ListenAndServe")
		}
	}()

	err = app.Run()
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
