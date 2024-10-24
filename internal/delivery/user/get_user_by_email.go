package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (d *UsersDelivery) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	var req UsersLoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})

		return
	}

	sessionID, username, err := d.userManager.GetUserByEmail(r.Context(), req.ToModel())
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.WrongCredentials.Error(),
		})

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session-id",
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Value:    sessionID,
	})

	utils.WriteJSON(w, http.StatusOK, UsersUsernameResponse{
		Username: username,
	})
}
