package main

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"

	"2024_2_kotyari/config"
	"2024_2_kotyari/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// @title Swagger Oxic API
// @description This is simple oxic server

// @host oxic.swagger.io
// @BasePath /
func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
		return
	}

	config.Init()
	server := handlers.NewServer(&config.Cfg)

	r := mux.NewRouter()

	r.HandleFunc("/login", server.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", server.LogoutHandler).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("Сервер запущен на", config.GetServerAddress())
	log.Fatal(http.ListenAndServe(config.GetServerAddress(), r))
}
