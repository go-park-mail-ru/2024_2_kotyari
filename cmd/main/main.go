package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"

	mainService "github.com/go-park-mail-ru/2024_2_kotyari/internal/apps/main_service"
	"github.com/joho/godotenv"
)

const configFile = ".env"

// @title Swagger Oxic API
// @version 1.0
// @description This is simple oxic server
// @host 94.139.246.241:8000
// @BasePath /
func main() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server, err := mainService.NewServer()
	if err != nil {
		log.Fatal(fmt.Errorf("error occured when creating server, %w", err))
	}

	router := mux.NewRouter()
	router.PathPrefix("/metrics").Handler(promhttp.Handler())
	serverProm := http.Server{Handler: router, Addr: fmt.Sprintf(":%d", 8080), ReadHeaderTimeout: 10 * time.Second}

	go func() {
		if err = serverProm.ListenAndServe(); err != nil {
			log.Println("fail auth.ListenAndServe")
		}
	}()

	if err = server.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
