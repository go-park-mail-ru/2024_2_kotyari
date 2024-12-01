package address

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (a *AddressAdapter) FetchAddress(ctx context.Context, address model.Addresses) (model.Addresses, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		a.log.Error("[AddressAdapter.FetchAddress] Failed to get requestID", slog.String("error", err.Error()))

		return model.Addresses{}, err
	}

	a.log.Info("[AddressAdapter.FetchAddress] Started executing", slog.Any("request-id", requestID))

	var addressAdapterResponseDTO AddressAdatperResponseDTO

	resp, err := a.HttpClient.Get(fmt.Sprintf("%s&text=%s", a.BaseURL, url.QueryEscape(address.Address)))
	if err != nil {
		a.log.Error("[AddressAdapter.FetchAddress] Failed to fetch api results", slog.String("error", err.Error()))

		return model.Addresses{}, err
	}

	fmt.Println("ААААААА: ", address, "aaaaa", resp)

	if err = json.NewDecoder(resp.Body).Decode(&addressAdapterResponseDTO); err != nil {
		a.log.Error("[AddressAdapter.FetchAddress] Invalid request body", slog.String("error", err.Error()))

		return model.Addresses{}, err
	}

	if len(addressAdapterResponseDTO.Results) == 0 {
		a.log.Error("[AddressAdapter.FetchAddress] No addresses to suggest")

		return model.Addresses{}, errs.NoAddressesToSuggest
	}

	return singleResultToModel(addressAdapterResponseDTO.Results[0].Title.Text), nil
}
