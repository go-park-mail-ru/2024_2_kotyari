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

	_, err := pd.client.ChangePassword(r.Context(), &profilegrpc.ChangePasswordRequest{
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
