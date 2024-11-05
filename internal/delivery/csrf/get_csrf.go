package csrf

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"net/http"
	"time"
)

func (csrf *CsrfDelivery) GetCsrf(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.SessionNotFound)

		return
	}

	cookie, _ := r.Cookie(utils.SessionName)
	session, _ := csrf.sessionGetter.Get(r.Context(), cookie.Value)

	token, err := csrf.csrfCreator.CreateCsrfToken(session, time.Now())
	if err != nil {
		return
	}

	w.Header().Set("X-CSRF-Token", token)
	w.WriteHeader(http.StatusOK)
}
