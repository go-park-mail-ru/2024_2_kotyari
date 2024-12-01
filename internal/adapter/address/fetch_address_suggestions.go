package address

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (a *AddressAdapter) FetchAddressSuggestions(ctx context.Context, q string) ([]model.Addresses, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		a.log.Error("[AddressAdapter.FetchAddressSuggestions] Failed to get requestID", slog.String("error", err.Error()))

		return nil, err
	}

	a.log.Info("[AddressAdapter.FetchAddressSuggestions] Started executing", slog.Any("request-id", requestID))

	var addressAdapterResponseDTO AddressAdatperResponseDTO

	resp, err := a.HttpClient.Get(fmt.Sprintf("%s&text=%s", a.BaseURL, q))
	if err != nil {
		a.log.Error("[AddressAdapter.FetchAddressSuggestions] Failed to fetch api results", slog.String("error", err.Error()))

		return nil, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&addressAdapterResponseDTO); err != nil {
		a.log.Error("[AddressAdapter.FetchAddressSuggestions] Invalid request body", slog.String("error", err.Error()))

		return nil, err
	}

	if len(addressAdapterResponseDTO.Results) == 0 {
		a.log.Error("[AddressAdapter.FetchAddressSuggestions] No addresses to suggest")

		return nil, errs.NoAddressesToSuggest
	}

	return addressAdapterResponseDTO.ToModelSlice(), nil
}
