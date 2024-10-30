package profile

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/address"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *ProfilesDelivery) GetProfile(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	h.log.Info("Обработка запроса на получение профиля", "request", request)

	id := utils.GetContextSessionUserID(request.Context())

	profile, err := h.profileManager.GetProfile(ctx, uint32(id))
	if err != nil {
		h.log.Error("Ошибка при получении профиля на уровне деливери", slog.String("error", err.Error()))
		return
	}

	addressResponse := address.AddressDTO{
		ID:        profile.Address.Id,
		City:      profile.Address.City,
		Street:    profile.Address.Street,
		House:     profile.Address.House,
		Flat:      profile.Address.Flat,
		ProfileID: id,
	}

	profileResponse := ProfileResponse{
		ID:        profile.ID,
		Email:     profile.Email,
		Username:  profile.Username,
		Age:       profile.Age,
		Address:   addressResponse,
		Gender:    profile.Gender,
		AvatarUrl: profile.AvatarUrl,
	}

	h.log.Info("Успешное получение профиля", profileResponse)

	utils.WriteJSON(writer, http.StatusOK, profileResponse)
}
