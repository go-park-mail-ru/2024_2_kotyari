package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (d *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req UsersSignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})

		return
	}

	req.Username = d.stringSanitizer.SanitizeString(req.Username)
	req.Email = d.stringSanitizer.SanitizeString(req.Email)
	req.Password = d.stringSanitizer.SanitizeString(req.Password)
	req.RepeatPassword = d.stringSanitizer.SanitizeString(req.RepeatPassword)

	if err = utils.ValidateRegistration(req.Email, req.Username, req.Password, req.RepeatPassword); err != nil {
		err, code := d.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	sessionID, user, err := d.userManager.CreateUser(r.Context(), req.ToModel())
	if err != nil {
		err, code := d.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	http.SetCookie(w, utils.SetSessionCookie(sessionID))

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		Username:  user.Username,
		City:      user.City,
		AvatarUrl: user.AvatarUrl,
	})
}
