package user

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
	"net/http"
)

func (u *UsersDelivery) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req UsersLoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})
		u.log.Error("[ UsersDelivery.LoginUser ] Ошибка при декодировании запроса", slog.String("error", err.Error()))
		return
	}

	usersDefaultResponse, err := u.userClientGrpc.LoginUser(r.Context(), req.ToGrpcLoginRequest())

	if err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
		u.log.Error("[ UsersDelivery.LoginUser ] Ошибка при отправке на grpc", slog.String("error", err.Error()))
		return
	}

	sessionID, err := u.sessionService.Create(r.Context(), usersDefaultResponse.UserId)
	if err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
		u.log.Error("[ UsersDelivery.LoginUser ] Ошибка при получении сессии", slog.String("error", err.Error()))
		return
	}

	http.SetCookie(w, utils.SetSessionCookie(sessionID))

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		UserID:   usersDefaultResponse.UserId,
		Username: usersDefaultResponse.Username,
		City:     usersDefaultResponse.City,
	})
}
