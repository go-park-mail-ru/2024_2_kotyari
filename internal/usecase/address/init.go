package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type addressRepository interface {
	CreateAddress(ctx context.Context, profileID uint32, addressModel model.Address) (uint32, error)
	GetAddressByProfileID(ctx context.Context, profileID uint32) (model.Address, error)
	UpdateUsersAddress(addressID uint32, addressModel model.Address) error
}

type AddressService struct {
	addressRepo addressRepository
	log         *slog.Logger
}

func NewAddressService(addressRepository addressRepository, logger *slog.Logger) *AddressService {
	return &AddressService{
		addressRepo: addressRepository,
		log:         logger,
	}
}
