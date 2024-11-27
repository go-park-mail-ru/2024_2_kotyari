package profile

import (
	"encoding/json"
	"log/slog"
	"net/http"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
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

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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
	if err != nil {
		err, i := pd.errResolver.Get(err)
		utils.WriteErrorJSON(w, i, err)

		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
