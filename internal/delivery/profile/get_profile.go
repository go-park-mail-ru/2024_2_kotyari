package profile

import (
	profile_grpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (pd *ProfilesDelivery) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	profile, err := pd.client.GetProfile(r.Context(), &profile_grpc.GetProfileRequest{UserId: userID})
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
