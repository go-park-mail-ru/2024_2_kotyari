package address

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *AddressDelivery) GetAddress(writer http.ResponseWriter, request *http.Request) {
	userID, ok := utils.GetContextSessionUserID(request.Context())
	if !ok {
		utils.WriteErrorJSON(writer, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	address, err := h.addressManager.GetAddressByProfileID(request.Context(), userID)

	if err != nil {
		h.log.Error("[ AddressDelivery.GetAddress ] Ошибка при получении адреса на уровне деливери", slog.String("error", err.Error()))
		return
	}

	addressResponse := FromModel(address)

	utils.WriteJSON(writer, http.StatusOK, addressResponse)
}
