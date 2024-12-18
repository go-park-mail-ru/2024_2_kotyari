package user

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/mailru/easyjson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UsersDelivery) LoginUser(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		u.log.Error("[UsersDelivery.LoginUser] No request ID")
		utils.WriteErrorJSONByError(w, err, u.errResolver)

		return
	}

	u.log.Info("[UsersDelivery.LoginUser] Started executing", slog.Any("request-id", requestID))

	var req UsersLoginRequest

	err = easyjson.UnmarshalFromReader(r.Body, &req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidJSONFormat.Error(),
		})
		u.log.Error("[ UsersDelivery.LoginUser ] Ошибка при декодировании запроса", slog.String("error", err.Error()))
		return
	}

	newCtx, err := utils.AddMetadataRequestID(r.Context())
	if err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	usersDefaultResponse, err := u.userClientGrpc.LoginUser(newCtx, req.ToGrpcLoginRequest())

	if err != nil {
		u.log.Error("[ UsersDelivery.LoginUser ] До вывода кода", slog.String("error", err.Error()), slog.String("func", "u.userClientGrpc.LoginUser"))

		grpcErr, ok := status.FromError(err)
		if ok {
			u.log.Error("[ UsersDelivery.LoginUser ] ",
				slog.String("error", err.Error()), slog.String("func", "u.userClientGrpc.LoginUser"),
				slog.Any("code", grpcErr.Code()))

			switch grpcErr.Code() {
			case codes.NotFound:
				err, code := u.errResolver.Get(errs.UserDoesNotExist)
				utils.WriteJSON(w, code, errs.HTTPErrorResponse{
					ErrorMessage: err.Error(),
				})
				return

			case codes.InvalidArgument:
				err, code := u.errResolver.Get(errs.WrongCredentials)
				utils.WriteJSON(w, code, errs.HTTPErrorResponse{
					ErrorMessage: err.Error(),
				})
				return

			default:
				err, code := u.errResolver.Get(errs.InternalServerError)
				utils.WriteJSON(w, code, errs.HTTPErrorResponse{
					ErrorMessage: err.Error(),
				})
				return
			}
		}

		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.InternalServerError.Error(),
		})
		u.log.Error("[ UsersDelivery.LoginUser ] ", slog.String("error", err.Error()), slog.String("func", "u.userClientGrpc.LoginUser"), slog.Int("code", code))
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
