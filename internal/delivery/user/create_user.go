package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UsersDelivery) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		u.log.Error("[UsersDelivery.CreateUser] No request ID")
		utils.WriteErrorJSONByError(w, err, u.errResolver)

		return
	}

	u.log.Info("[UsersDelivery.CreateUser] Started executing", slog.Any("request-id", requestID))

	var req UsersSignUpRequest

	err = json.NewDecoder(r.Body).Decode(&req)
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
	newCtx, err := utils.AddMetadataRequestID(r.Context())
	if err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	usersDefaultResponse, err := u.userClientGrpc.CreateUser(newCtx, req.ToGrpcSignupRequest())

	grpcErr, ok := status.FromError(err)
	if err != nil {
		if ok {
			switch grpcErr.Code() {
			case codes.InvalidArgument:
				u.log.Error("[ UsersDelivery.CreateUser ] Пользователь уже существует", err.Error())

				utils.WriteErrorJSONByError(w, errs.UserAlreadyExists, u.errResolver)

				return

			default:
				u.log.Error("[ UsersDelivery.CreateUser ] Неизвестная ошибка", err.Error())

				utils.WriteErrorJSONByError(w, errs.InternalServerError, u.errResolver)

				return
			}
		}

		u.log.Error("[ UsersDelivery.CreateUser ] Не удалось получить код ошибки")

		utils.WriteErrorJSONByError(w, errs.InternalServerError, u.errResolver)

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
