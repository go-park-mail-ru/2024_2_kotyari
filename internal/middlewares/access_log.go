package middlewares

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
	"net/http"
	"time"
)

type responseStatusWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseStatusWrapper) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func AccessLogMiddleware(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startRequestTime := time.Now()

			responseWriter := &responseStatusWrapper{
				w,
				http.StatusOK,
			}

			next.ServeHTTP(responseWriter, r)
			requestDuration := time.Since(startRequestTime)
			requestID, err := utils.GetContextRequestID(r.Context())
			if err != nil {
				log.Error("[AccessLogMiddleware] Empty requestID")
			}

			log.Info("Access Log",
				"method", r.Method,
				"url", r.URL.String(),
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
				"status", responseWriter.statusCode,
				"duration_ms", requestDuration.Milliseconds(),
				"request_id", requestID,
			)
		})
	}
}
