package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type addressManager interface {
	CreateUsersAddress(ctx context.Context, userID uint32, addressInfo model.Address) (uint32, error)
	GetAddressByProfileID(ctx context.Context, userID uint32) (model.Address, error)
	UpdateUsersAddress(addressID uint32, newAddress model.Address) error
}

type AddressDelivery struct {
	addressManager addressManager
	log            *slog.Logger
}

func NewAddressHandler(addressManager addressManager, logger *slog.Logger) *AddressDelivery {
	return &AddressDelivery{
		addressManager: addressManager,
		log:            logger,
	}
}
