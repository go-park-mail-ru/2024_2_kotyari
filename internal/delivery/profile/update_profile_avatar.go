package profile

import (
	"errors"
	"log/slog"
	"net/http"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

const maxUploadSize = 1024 * 1024 * 10

func (pd *ProfilesDelivery) UpdateProfileAvatar(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		pd.log.Error("[ProfilesDelivery.UpdateProfileAvatar] No request ID")
		utils.WriteErrorJSONByError(w, err, pd.errResolver)

		return
	}

	pd.log.Info("[ProfilesDelivery.UpdateProfileAvatar] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	avatarPath, msg, err := pd.uploadAvatarFromRequest(r)
	if err != nil {
		pd.log.Error("UpdateProfileAvatar",
			slog.String("error", err.Error()),
			"user", userID,
		)

		utils.WriteErrorJSON(w, http.StatusInternalServerError, msg)

		return
	}

	newCtx, err := utils.AddMetadataRequestID(r.Context())
	if err != nil {
		err, code := pd.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	_, err = pd.client.UpdateProfileAvatar(
		newCtx,
		&profilegrpc.UpdateAvatarRequest{
			UserId:   userID,
			Filepath: avatarPath})
	if err != nil {
		pd.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ]", slog.String("error", err.Error()))
		if errors.Is(err, errs.ErrFileTypeNotAllowed) {
			utils.WriteErrorJSON(w, http.StatusBadRequest, errs.ErrFileTypeNotAllowed)

			return
		}

		utils.WriteErrorJSON(w, http.StatusInternalServerError, errors.New("не удалось обновить аватар профиля, попробуйте позже"))

		return
	}

	res := AvatarResponse{AvatarUrl: avatarPath}

	utils.WriteJSON(w, http.StatusOK, res)
}
