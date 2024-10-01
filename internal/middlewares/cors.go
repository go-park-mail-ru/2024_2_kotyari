package middlewares

import (
	"net/http"
)

func setupCORS(w http.ResponseWriter, req *http.Request, sessionLifetime string) {
	allowedOrigins := map[string]struct{}{
		"http://localhost:3000":      {},
		"http://localhost:8080":      {},
		"http://127.0.0.1:3000":      {},
		"http://127.0.0.1:8080":      {},
		"94.139.246.241":             {},
		"http://94.139.246.241:8000": {},
		"http://94.139.246.241":      {},
	}

	origin := req.Header.Get("Origin")
	if _, ok := allowedOrigins[origin]; ok {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	w.Header().Set("Access-Control-Allow-Methods", http.MethodPost+", "+http.MethodGet+", "+http.MethodOptions+", "+http.MethodPut+", "+http.MethodDelete)
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Authorization, Access-Control-Allow-Origin")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Max-Age", sessionLifetime)
}

func CorsMiddleware(next http.Handler, sessionLifetime string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		setupCORS(w, req, sessionLifetime)

		// Если метод OPTIONS, то возвращаем пустой ответ с нужными заголовками
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, req)
	})
}
