package user

import (
	"encoding/json"
	grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"net/http"
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

	if err = utils.ValidateRegistration(req.Email, req.Username, req.Password, req.RepeatPassword); err != nil {
		err, code := d.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}
	usersDefaultResponse, err := d.userClientGrpc.CreateUser(r.Context(), &grpc_gen.UsersSignUpRequest{
		Username:       req.Username,
		Email:          req.Email,
		Password:       req.Password,
		RepeatPassword: req.RepeatPassword,
	})
	if err != nil {
		err, code := d.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	sessionID, err := d.sessionService.Create(r.Context(), usersDefaultResponse.UserId)
	if err != nil {
		err, code := d.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		return
	}

	http.SetCookie(w, utils.SetSessionCookie(sessionID))

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		UserID:   usersDefaultResponse.UserId,
		Username: usersDefaultResponse.Username,
		City:     usersDefaultResponse.City,
	})
}
