package sessions

import (
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (sd *SessionDelivery) Delete(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(utils.SessionName)
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

	err = sd.sessionRemover.Delete(r.Context(), model.Session{SessionID: cookie.Value})
	if err != nil {
		err, code := sd.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	http.SetCookie(w, utils.RemoveSessionCookie())

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
