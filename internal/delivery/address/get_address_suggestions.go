package address

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (h *AddressDelivery) GetAddressSuggestions(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		h.log.Error("[AddressDelivery.GetAddressSuggestions] No request ID")
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	h.log.Info("[AddressDelivery.GetAddressSuggestions] Started executing", slog.Any("request-id", requestID))

	query := utils.GetSearchQuery(r)

	addresses, err := h.addressSuggestions.FetchAddressSuggestions(r.Context(), query)
	if err != nil {
		h.log.Error("[AddressDelivery.GetAddressSuggestions] Failed to fetch address suggestions",
			slog.String("error", err.Error()))
		utils.WriteErrorJSONByError(w, err, h.errResolver)

		return
	}

	addressSuggestionsResponse := make([]GetAddressResponse, 0, len(addresses))
	for _, address := range addresses {
		addressSuggestionsResponse = append(addressSuggestionsResponse, addressFromModel(address))
	}

	utils.WriteJSON(w, http.StatusOK, addressesSuggestionsFromSlice(addressSuggestionsResponse))
}
