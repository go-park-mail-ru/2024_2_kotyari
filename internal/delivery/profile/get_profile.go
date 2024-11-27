package profile

import (
	"log/slog"
	"net/http"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (pd *ProfilesDelivery) GetProfile(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		pd.log.Error("[ProfilesDelivery.GetProfile] No request ID")
		utils.WriteErrorJSONByError(w, err, pd.errResolver)

		return
	}

	pd.log.Info("[ProfilesDelivery.GetProfile] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	newCtx, err := utils.AddMetadataRequestID(r.Context())
	if err != nil {
		err, code := pd.errResolver.Get(err)
		utils.WriteJSON(w, code, errs.HTTPErrorResponse{
			ErrorMessage: err.Error(),
		})
	}

	profile, err := pd.client.GetProfile(newCtx, &profilegrpc.GetProfileRequest{UserId: userID})
	if err != nil {
		pd.log.Error("[ ProfilesDelivery.GetProfile ]",
			slog.String("error", err.Error()),
		)

		_, code := pd.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)
		return
	}

	address, err := pd.addressGetter.GetAddressByProfileID(r.Context(), profile.UserId)
	if err != nil {
		pd.log.Error("[ ProfilesDelivery.GetProfile ]",
			slog.String("error", err.Error()),
		)

		_, code := pd.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	profileResponse := fromGrpcResponse(profile, address)

	utils.WriteJSON(w, http.StatusOK, profileResponse)
}
