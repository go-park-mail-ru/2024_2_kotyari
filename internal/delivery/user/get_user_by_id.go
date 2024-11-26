package user

import (
	"log/slog"
	"net/http"

	grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (u *UsersDelivery) GetUserById(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		u.log.Error("[UsersDelivery.GetUserById] No request ID")
		utils.WriteErrorJSONByError(w, err, u.errResolver)

		return
	}

	u.log.Info("[UsersDelivery.GetUserById] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
		u.log.Error("[ UsersDelivery.GetUserById ] Пользователь не авторизован")
		return
	}

	newCtx, err := utils.AddMetadataRequestID(r.Context())
	if err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	usersDefaultResponse, err := u.userClientGrpc.GetUserById(newCtx, &grpc_gen.GetUserByIdRequest{UserId: userID})
	if err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
		u.log.Error("[ UsersDelivery.GetUserById ] Ошибка при отправке на grpc", slog.String("error", err.Error()))
		return
	}

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		Username:  usersDefaultResponse.Username,
		City:      usersDefaultResponse.City,
		AvatarUrl: usersDefaultResponse.AvatarUrl,
	})
}
