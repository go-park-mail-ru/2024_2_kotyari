package middlewares

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/auth"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func AuthMiddleware(sessions auth.SessionInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := sessions.Get(r)
			if err != nil {
				utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
					ErrorMessage: errs.InternalServerError.Error(),
				})

				return
			}
			email, exists := session.Values[auth.SessionKey].(string)
			if !exists {
				utils.WriteJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
					ErrorMessage: errs.UserNotAuthorized.Error(),
				})

				return
			}

			ctx := context.WithValue(r.Context(), auth.SessionKey, email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
