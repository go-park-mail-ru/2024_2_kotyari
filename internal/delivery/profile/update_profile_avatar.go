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

func (pd *ProfilesDelivery) UpdateProfileAvatar(writer http.ResponseWriter, request *http.Request) {
	userID, ok := utils.GetContextSessionUserID(request.Context())
	if !ok {
		utils.WriteErrorJSON(writer, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	avatarPath, msg, err := pd.uploadAvatarFromRequest(request)
	if err != nil {
		pd.log.Error("UpdateProfileAvatar",
			slog.String("error", err.Error()),
			"user", userID,
		)

		utils.WriteErrorJSON(writer, http.StatusInternalServerError, msg)

		return
	}

	resGrpc, err := pd.client.UpdateProfileAvatar(
		request.Context(),
		&profilegrpc.UpdateAvatarRequest{
			UserId:   userID,
			Filepath: avatarPath,
		})
	if err != nil {
		pd.log.Error("[ ProfilesDelivery.UpdateProfileAvatar ]", slog.String("error", err.Error()))
		if errors.Is(err, errs.ErrFileTypeNotAllowed) {
			utils.WriteErrorJSON(writer, http.StatusBadRequest, errs.ErrFileTypeNotAllowed)

			return
		}

		utils.WriteErrorJSON(writer, http.StatusInternalServerError, errors.New("не удалось обновить аватар профиля, попробуйте позже"))

		return
	}

	avatarResponse := AvatarResponse{
		AvatarUrl: resGrpc.Filepath,
	}

	utils.WriteJSON(writer, http.StatusOK, avatarResponse)
}
