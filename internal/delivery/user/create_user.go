package user

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	grpc_gen "path/to/your/protobuf/generated/package"
)

func (d *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req grpc_gen.UsersSignUpRequest
	if err := utils.DecodeJSONBody(r, &req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})

		return
	}

	// Валидация и очистка входных данных
	req.Username = d.stringSanitizer.SanitizeString(req.GetUsername())
	req.Email = d.stringSanitizer.SanitizeString(req.GetEmail())
	req.Password = d.stringSanitizer.SanitizeString(req.GetPassword())
	req.RepeatPassword = d.stringSanitizer.SanitizeString(req.GetRepeatPassword())

	if err := utils.ValidateRegistration(req.Email, req.Username, req.Password, req.RepeatPassword); err != nil {
		httpError, code := d.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: httpError.Error(),
		})
		return
	}

	// Вызов gRPC метода CreateUser
	ctx := r.Context()
	response, err := d.grpcClient.CreateUser(ctx, &req)
	if err != nil {
		httpError, code := d.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: httpError.Error(),
		})
		return
	}

	// Создание cookie с сессией
	sessionCookie := utils.SetSessionCookie(response.GetUsername()) // предполагается, что идентификатор сессии возвращается в gRPC
	http.SetCookie(w, sessionCookie)

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		Username:  user.Username,
		City:      user.City,
		AvatarUrl: user.AvatarUrl,
	})
}
