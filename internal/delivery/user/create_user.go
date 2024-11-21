package user

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
	"net/http"
)

func (u *UsersDelivery) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req UsersSignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})
		u.log.Error("[ UsersDelivery.CreateUser ] Ошибка при декодировании запроса", slog.String("error", err.Error()))
		return
	}

	if err = utils.ValidateRegistration(req.Email, req.Username, req.Password, req.RepeatPassword); err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
		u.log.Error("[ UsersDelivery.CreateUser ] Валидация регистрации не прошла успешно", slog.String("error", err.Error()))
		return
	}
	usersDefaultResponse, err := u.userClientGrpc.CreateUser(r.Context(), req.ToGrpcSignupRequest())
	if err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		u.log.Error("[ UsersDelivery.CreateUser ] Ошибка при передаче на grpc", slog.String("error", err.Error()))
		return
	}

	sessionID, err := u.sessionService.Create(r.Context(), usersDefaultResponse.GetUserId())
	if err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})

		u.log.Error("[ UsersDelivery.CreateUser ] Ошибка при создании сессии ",
			slog.String("error", err.Error()),
			slog.Any("userId", usersDefaultResponse.UserId),
		)

		return
	}

	http.SetCookie(w, utils.SetSessionCookie(sessionID))

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		UserID:   usersDefaultResponse.UserId,
		Username: usersDefaultResponse.Username,
		City:     usersDefaultResponse.City,
	})
}
