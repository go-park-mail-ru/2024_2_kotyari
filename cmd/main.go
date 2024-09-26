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

// @title Product Catalog API
// @version 1.0
// @description API для получения продуктов из каталога.
// @host localhost:8000
// @BasePath /catalog
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
		return
	}

	config.Init()
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/catalog/products", handlers.ProductsHandler).Methods("GET")
	r.HandleFunc("/catalog/product/{id}", handlers.ProductByIDHandler).Methods("GET")

	fmt.Println("Сервер запущен на", config.GetServerAddress())
	log.Fatal(http.ListenAndServe(config.GetServerAddress(), r))
}
