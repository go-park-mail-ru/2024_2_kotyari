package address

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type addressManager interface {
	GetAddressByProfileID(ctx context.Context, userID uint32) (model.Address, error)
	UpdateUsersAddress(ctx context.Context, addressID uint32, newAddress model.Address) error
}

type AddressDelivery struct {
	addressManager addressManager
	errResolver    errs.GetErrorCode
	log            *slog.Logger
}

func NewAddressHandler(addressManager addressManager, errResolver errs.GetErrorCode, logger *slog.Logger) *AddressDelivery {
	return &AddressDelivery{
		addressManager: addressManager,
		errResolver:    errResolver,
		log:            logger,
	}
}
