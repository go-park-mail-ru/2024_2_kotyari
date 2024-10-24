package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (d *UsersDelivery) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req UsersSignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})

		return
	}

	if err = utils.ValidateRegistration(req.Email, req.Username, req.Password, req.RepeatPassword); err != nil {
		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	sessionID, user, err := d.userManager.CreateUser(r.Context(), req.ToModel())
	if err != nil {
		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	/// TODO: Remove magic constant
	http.SetCookie(w, &http.Cookie{
		Name:     "session-id",
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Value:    sessionID,
	})

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		Username: user.Username,
		City:     user.City,
	})
}
