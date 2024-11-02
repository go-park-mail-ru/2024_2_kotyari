package address

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *AddressDelivery) GetAddress(writer http.ResponseWriter, request *http.Request) {

	id := utils.GetContextSessionUserID(request.Context())

	address, err := h.addressManager.GetAddressByProfileID(request.Context(), uint32(id))

	if err != nil {
		h.log.Error("[ AddressDelivery.GetAddress ] Ошибка при получении адреса на уровне деливери", slog.String("error", err.Error()))
		return
	}

	addressResponse := FromModel(address)

	utils.WriteJSON(writer, http.StatusOK, addressResponse)
}
