package user

import (
	grpc_gen "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (u *UsersDelivery) GetUserById(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
		u.log.Error("[ UsersDelivery.GetUserById ] Пользователь не авторизован")
		return
	}
	usersDefaultResponse, err := u.userClientGrpc.GetUserById(r.Context(), &grpc_gen.GetUserByIdRequest{UserId: userID})
	if err != nil {
		err, code := u.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
		u.log.Error("[ UsersDelivery.GetUserById ] Ошибка при отправке на grpc", slog.String("error", err.Error()))
		return
	}

	utils.WriteJSON(w, http.StatusOK, UsersDefaultResponse{
		Username: usersDefaultResponse.Username,
		City:     usersDefaultResponse.City,
	})
}
