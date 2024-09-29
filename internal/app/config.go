package app

import (
	"github.com/rs/cors"
	"net/http"
)

const (
	second = 1
	minute = 60 * second
	hour   = 60 * minute
)

func setUpCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://localhost:8080",
			"http://127.0.0.1:3000",
			"http://127.0.0.1:8080",
			"94.139.246.241",
			"94.139.246.241:8000",
			"http://94.139.246.241",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions},
		AllowedHeaders: []string{
			"Accept",
			"Accept-Language",
			"Content-Type",
			"Authorization",
			"Access-Control-Allow-Origin",
		},
		AllowCredentials: true,
		MaxAge:           hour,
		Debug:            false,
	})
}
