package app

import (
	"net/http"
	"strconv"
)

const (
	second = 1
	minute = 60 * second
	hour   = 60 * minute
)

func setupCORS(w http.ResponseWriter, req *http.Request) {
	allowedOrigins := map[string]bool{
		"http://localhost:3000":      true,
		"http://localhost:8080":      true,
		"http://127.0.0.1:3000":      true,
		"http://127.0.0.1:8080":      true,
		"94.139.246.241":             true,
		"http://94.139.246.241:8000": true,
		"http://94.139.246.241":      true,
	}

	if origin := req.Header.Get("Origin"); allowedOrigins[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Authorization, Access-Control-Allow-Origin")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Max-Age", strconv.Itoa(hour))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		setupCORS(w, req)

		// Если метод OPTIONS, то возвращаем пустой ответ с нужными заголовками
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, req)
	})
}
