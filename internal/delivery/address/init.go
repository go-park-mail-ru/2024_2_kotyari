package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type addressSuggestionsFetcher interface {
	FetchAddressSuggestions(ctx context.Context, q string) ([]model.Addresses, error)
}

type addressManager interface {
	GetAddressByProfileID(ctx context.Context, userID uint32) (model.Addresses, error)
	UpdateUsersAddress(ctx context.Context, userID uint32, newAddress model.Addresses) error
}

type AddressDelivery struct {
	addressManager     addressManager
	addressSuggestions addressSuggestionsFetcher
	errResolver        errs.GetErrorCode
	stringSanitizer    utils.StringSanitizer
	log                *slog.Logger
}

func NewAddressHandler(addressManager addressManager, addressSuggestions addressSuggestionsFetcher,
	errResolver errs.GetErrorCode, stringSanitizer utils.StringSanitizer, logger *slog.Logger) *AddressDelivery {
	return &AddressDelivery{
		addressManager:     addressManager,
		addressSuggestions: addressSuggestions,
		errResolver:        errResolver,
		stringSanitizer:    stringSanitizer,
		log:                logger,
	}
}
