package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type addressFetcher interface {
	FetchAddress(ctx context.Context, address model.Addresses) (model.Addresses, error)
}

type addressRepository interface {
	GetAddressByProfileID(ctx context.Context, profileID uint32) (model.Addresses, error)
	UpdateUsersAddress(ctx context.Context, addressID uint32, addressModel model.Addresses) error
}

type AddressService struct {
	addressRepo    addressRepository
	addressFetcher addressFetcher
	log            *slog.Logger
}

func NewAddressService(addressRepository addressRepository, addressFetcher addressFetcher, logger *slog.Logger) *AddressService {
	return &AddressService{
		addressRepo:    addressRepository,
		addressFetcher: addressFetcher,
		log:            logger,
	}
}
