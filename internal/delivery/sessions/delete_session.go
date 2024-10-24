package sessions

import (
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (sd *SessionDelivery) Delete(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-id")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
				ErrorMessage: err.Error(),
			})

			return
		}

		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.InternalServerError.Error(),
		})

		return
	}

	err = sd.sessionRemover.Delete(r.Context(), model.Session{SessionID: cookie.Value})
	if err != nil {
		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session-id",
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
