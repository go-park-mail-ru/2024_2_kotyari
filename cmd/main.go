package main

import (
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"2024_2_kotyari/config"
	"2024_2_kotyari/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// @title Swagger Oxic API
// @version 1.0
// @description This is simple oxic server

// @host oxic.swagger.io
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
		return
	}

	config.Init()
	server := handlers.NewServer(&config.Cfg)

	r := mux.NewRouter()

	r.HandleFunc("/login", server.Login).Methods(http.MethodPost)
	r.HandleFunc("/logout", server.Logout).Methods(http.MethodPost)
	r.HandleFunc("/signup", server.Signup).Methods(http.MethodPost)
	r.HandleFunc("/catalog/products", handlers.Products).Methods("GET")
	r.HandleFunc("/catalog/product/{id}", handlers.ProductByID).Methods("GET")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("Сервер запущен на", config.GetServerAddress())
	log.Fatal(http.ListenAndServe(config.GetServerAddress(), r))
}
