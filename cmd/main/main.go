package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/apps/main_service"
	"log"

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

	server, err := main_service.NewServer()
	if err != nil {
		log.Fatal(fmt.Errorf("error occured when creating server, %w", err))
	}

	if err := server.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
