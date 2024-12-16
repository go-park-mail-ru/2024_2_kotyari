package profile

import (
	"encoding/json"
	profile_grpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (pd *ProfilesDelivery) UpdateProfileData(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		pd.log.Error("[ProfilesDelivery.UpdateProfileData] No request ID")
		utils.WriteErrorJSONByError(w, err, pd.errResolver)

		return
	}

	pd.log.Info("[ProfilesDelivery.UpdateProfileData] Started executing", slog.Any("request-id", requestID))

	userId, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}
	var req UpdateProfile

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		pd.log.Error("[ ProfilesDelivery.UpdateProfileData ] Ошибка десериализации запроса",
			slog.String("error", err.Error()),
		)

		utils.WriteJSON(w, http.StatusBadRequest, &errs.HTTPErrorResponse{ErrorMessage: errs.InvalidJSONFormat.Error()})

		return
	}

	newCtx, err := utils.AddMetadataRequestID(r.Context())
	if err != nil {
		err, code := pd.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	_, err = pd.client.UpdateProfileData(newCtx, &profile_grpc.UpdateProfileDataRequest{
		UserId:   userId,
		Email:    req.Email,
		Username: req.Username,
		Gender:   req.Gender})
	if err != nil {
		_, code := pd.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	res := UpdateProfile{
		Email:    req.Email,
		Username: req.Username,
		Gender:   req.Gender,
	}

	utils.WriteJSON(w, http.StatusOK, res)
}
