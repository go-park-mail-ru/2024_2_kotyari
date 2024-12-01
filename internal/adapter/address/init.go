package address

import (
	"fmt"
	"log/slog"
	"net/http"
)

type AddressAdapter struct {
	BaseURL    string
	log        *slog.Logger
	HttpClient *http.Client
}

func NewAddressAdapter(apiKey string, logger *slog.Logger) *AddressAdapter {
	baseURL := fmt.Sprintf("https://suggest-maps.yandex.ru/v1/suggest?apikey=%s", apiKey)

	return &AddressAdapter{
		BaseURL:    baseURL,
		log:        logger,
		HttpClient: &http.Client{},
	}
}
