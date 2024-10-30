package address

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *AddressDelivery) GetAddress(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	h.log.Info("Обработка запроса на получение адреса", "request", request)

	id := utils.GetContextSessionUserID(request.Context())

	address, err := h.addressManager.GetAddressByProfileID(ctx, uint32(id))

	if err != nil {
		h.log.Error("Ошибка при получении адреса на уровне деливери", slog.String("error", err.Error()))
		return
	}

	addressResponse := AddressDTO{
		ID:        address.Id,
		City:      address.City,
		Street:    address.Street,
		House:     address.House,
		Flat:      address.Flat,
		ProfileID: id,
	}

	h.log.Info("Успешное получение профиля", "id", id)

	utils.WriteJSON(writer, http.StatusOK, addressResponse)
}
