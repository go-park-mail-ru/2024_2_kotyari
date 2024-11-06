package middlewares

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New()

		ctx := context.WithValue(r.Context(), utils.RequestIDName, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
