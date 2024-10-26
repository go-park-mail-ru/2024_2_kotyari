package middlewares

import (
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie(utils.SessionName)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				utils.WriteJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
					ErrorMessage: errs.UserNotAuthorized.Error(),
				})

				return
			}

			utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
				ErrorMessage: errs.InternalServerError.Error(),
			})

			return
		}

		next.ServeHTTP(w, r)
	})
}
