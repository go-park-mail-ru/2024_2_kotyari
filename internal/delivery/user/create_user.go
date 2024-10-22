package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"net/http"
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
	//var resp CreateUserResponse

	username, err := d.userManager.CreateUser(r.Context(), req.ToModel())
	fmt.Println("Username: ", username)
	if err != nil {
		utils.WriteJSON(w, errs.ErrCodesMapping[err], errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	utils.WriteJSON(w, http.StatusOK, CreateUserResponse{
		Username: username,
	})
}
