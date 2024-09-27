package main

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/app"
	"github.com/joho/godotenv"
	"log"
)

// @title Swagger Oxic API
// @version 1.0
// @description This is simple oxic server

// @host localhost:8000
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.NewServer().Run()
}
