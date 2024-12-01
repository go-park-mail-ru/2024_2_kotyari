package address

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *AddressDelivery) GetAddress(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.log.Error("[AddressDelivery.GetAddress] No r ID")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.log.Info("[AddressDelivery.GetAddress] Started executing", slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)
	}

	address, err := h.addressManager.GetAddressByProfileID(r.Context(), userID)

	if err != nil {
		h.log.Error("[ AddressDelivery.GetAddress ] Ошибка при получении адреса на уровне деливери", slog.String("error", err.Error()))
		return
	}

	addressResponse := addressFromModel(address)

	utils.WriteJSON(w, http.StatusOK, addressResponse)
}
