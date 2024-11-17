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
			ErrorMessage: "Invalid request body",
		})
		return
	}

	// Валидация и очистка входных данных
	req.Username = d.inputValidator.SanitizeString(req.GetUsername())
	req.Email = d.inputValidator.SanitizeString(req.GetEmail())
	req.Password = d.inputValidator.SanitizeString(req.GetPassword())
	req.RepeatPassword = d.inputValidator.SanitizeString(req.GetRepeatPassword())

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

	// Успешный ответ
	utils.WriteJSON(w, http.StatusOK, response)
}
