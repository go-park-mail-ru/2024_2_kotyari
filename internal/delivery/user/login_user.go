package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (d *UsersHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req UsersLoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})

		return
	}

	req.Email = d.inputValidator.SanitizeString(req.Email)
	req.Password = d.inputValidator.SanitizeString(req.Password)

	sessionID, user, err := d.userClientGrpc.LoginUser(r.Context(), req.ToModel())
	if err != nil {
		err, code := d.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	http.SetCookie(w, utils.SetSessionCookie(sessionID))

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		Username: user.Username,
		City:     user.City,
	})
}
