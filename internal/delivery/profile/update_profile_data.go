package profile

import (
	"encoding/json"
	profile_grpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (pd *ProfilesDelivery) UpdateProfileData(writer http.ResponseWriter, request *http.Request) {
	userId, ok := utils.GetContextSessionUserID(request.Context())
	if !ok {
		utils.WriteErrorJSON(writer, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}
	var req UpdateProfile

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		pd.log.Error("[ ProfilesDelivery.UpdateProfileData ] Ошибка десериализации запроса",
			slog.String("error", err.Error()),
		)

		utils.WriteJSON(writer, http.StatusBadRequest, &errs.HTTPErrorResponse{ErrorMessage: errs.InvalidJSONFormat.Error()})

		return
	}

	_, err := pd.client.UpdateProfileData(request.Context(), &profile_grpc.UpdateProfileDataRequest{
		UserId:   userId,
		Email:    req.Email,
		Username: req.Username,
		Gender:   req.Gender})
	if err != nil {
		_, code := pd.errResolver.Get(err)
		utils.WriteErrorJSON(writer, code, err)

		return
	}

	res := UpdateProfile{
		Email:    req.Email,
		Username: req.Username,
		Gender:   req.Gender,
	}

	utils.WriteJSON(writer, http.StatusOK, res)
}
