package main

import (
	"fmt"
	"log"
	"net/http"

	"2024_2_kotyari/config"
	"2024_2_kotyari/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/protected", handlers.ProtectedHandler).Methods("GET")

	fmt.Println("Сервер запущен на", config.ServerAddress)
	log.Fatal(http.ListenAndServe(config.ServerAddress, r))
}
