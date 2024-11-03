package profile

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *ProfilesDelivery) GetProfile(writer http.ResponseWriter, request *http.Request) {

	id := utils.GetContextSessionUserID(request.Context())

	profile, err := h.profileManager.GetProfile(request.Context(), uint32(id))
	if err != nil {
		h.log.Error("[ ProfilesDelivery.GetProfile ] Ошибка при получении профиля на уровне деливери", slog.String("error", err.Error()))
		return
	}

	profileResponse := FromModel(profile)

	utils.WriteJSON(writer, http.StatusOK, profileResponse)
}
