package profile

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (pd *ProfilesDelivery) GetProfile(writer http.ResponseWriter, request *http.Request) {
	userID, ok := utils.GetContextSessionUserID(request.Context())
	if !ok {
		utils.WriteErrorJSON(writer, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	profile, err := pd.profileManager.GetProfile(request.Context(), userID)
	if err != nil {
		pd.log.Error("[ ProfilesDelivery.GetProfile ] Ошибка при получении профиля на уровне деливери", slog.String("error", err.Error()))
		return
	}

	profileResponse := FromModel(profile)

	utils.WriteJSON(writer, http.StatusOK, profileResponse)
}
