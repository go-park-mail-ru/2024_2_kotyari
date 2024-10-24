package user

import (
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (d *UsersDelivery) GetUserById(w http.ResponseWriter, r *http.Request) {
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

	username, city, err := d.userManager.GetUserBySessionID(r.Context(), cookie.Value)
	if err != nil {
		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		Username: username,
		City:     city,
	})
}
