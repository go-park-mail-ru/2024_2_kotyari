package metrics

import (
	"context"
	metrics "github.com/go-park-mail-ru/2024_2_kotyari/internal/metrics/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type responseStatusWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseStatusWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func CreateMetricsMiddleware(metric *metrics.HTTPMetrics) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			//request = request.WithContext(context.WithValue(request.Context(), "code", code))
			start := time.Now()

			responseWriter := &responseStatusWrapper{
				writer,
				http.StatusOK,
			}

			next.ServeHTTP(responseWriter, request)

			end := time.Since(start)

			request = request.WithContext(context.WithValue(request.Context(), "code", responseWriter.statusCode))

			codeStr := strconv.Itoa(responseWriter.statusCode)

			route := mux.CurrentRoute(request)
			path, _ := route.GetPathTemplate()

			log.Println(path, codeStr, end)

			if path != "/metrics" {
				metric.AddDuration(path, codeStr, end)
				metric.IncreaseTotal(path, codeStr)
			}
		})
	}
}
