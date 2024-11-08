package middlewares

import (
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"net/http"
)

func CSRFMiddleware(csrfValidator csrfValidator, sessionGetter sessionGetter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet || r.Method == http.MethodOptions {
				next.ServeHTTP(w, r)

				return
			}

			token := r.Header.Get("X-CSRF-Token")
			if token == "" {
				utils.WriteErrorJSON(w, http.StatusForbidden, errors.New("CSRF токен отсутствует"))

				return
			}

			_, ok := utils.GetContextSessionUserID(r.Context())
			if !ok {
				utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.SessionNotFound)

				return
			}

			cookie, _ := r.Cookie(utils.SessionName)
			session, _ := sessionGetter.Get(r.Context(), cookie.Value)

			valid, err := csrfValidator.IsValidCSRFToken(session, token)
			if err != nil || !valid {
				utils.WriteErrorJSON(w, http.StatusForbidden, errors.New("неверный CSRF токен"))

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
