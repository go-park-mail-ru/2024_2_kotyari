package profile

import (
	"log/slog"
	"net/http"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/mailru/easyjson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (pd *ProfilesDelivery) ChangePassword(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		pd.log.Error("[ProfilesDelivery.ChangePassword] No request ID")
		utils.WriteErrorJSONByError(w, err, pd.errResolver)

		return
	}

	pd.log.Info("[ProfilesDelivery.ChangePassword] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	var req UpdatePasswordRequest

	if err = easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
		pd.log.Error("[ ProfilesDelivery.UpdateProfileData ] Ошибка десериализации запроса",
			slog.String("error", err.Error()),
		)

		utils.WriteErrorJSON(w, http.StatusBadRequest, errs.InvalidJSONFormat)

		return
	}

	newCtx, err := utils.AddMetadataRequestID(r.Context())
	if err != nil {
		err, code := pd.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	_, err = pd.client.ChangePassword(newCtx, &profilegrpc.ChangePasswordRequest{
		UserId:         userID,
		OldPassword:    req.OldPassword,
		NewPassword:    req.NewPassword,
		RepeatPassword: req.RepeatPassword})

	grpcErr, ok := status.FromError(err)
	if err != nil {
		if ok {
			switch grpcErr.Code() {
			case codes.InvalidArgument:
				pd.log.Error("[ ProfilesDelivery.ChangePassword ] Неправильный пароль", grpcErr.String())

				utils.WriteErrorJSONByError(w, errs.WrongPassword, pd.errResolver)

				return

			case codes.Unauthenticated:
				pd.log.Error("[ ProfilesDelivery.ChangePassword ] Пользователь уже существует", grpcErr.String())

				utils.WriteErrorJSONByError(w, errs.PasswordsDoNotMatch, pd.errResolver)

				return

			default:
				pd.log.Error("[ ProfilesDelivery.ChangePassword ] Неизвестная ошибка", err.Error())

				utils.WriteErrorJSONByError(w, errs.InternalServerError, pd.errResolver)

				return
			}
		}

		pd.log.Error("[ UsersDelivery.CreateUser ] Не удалось получить код ошибки")

		utils.WriteErrorJSONByError(w, errs.InternalServerError, pd.errResolver)

		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
