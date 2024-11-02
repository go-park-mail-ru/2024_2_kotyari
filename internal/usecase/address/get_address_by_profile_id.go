package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (as *AddressService) GetAddressByProfileID(ctx context.Context, userID uint32) (model.AddressDTO, error) {

	address, err := as.addressRepo.GetAddressByProfileID(ctx, userID)
	if err != nil {
		as.log.Error("[ AddressService.GetAddressByProfileID ] Ошибка при получении адреса", slog.String("error", err.Error()))
		return model.AddressDTO{}, err
	}
	if address.Flat == nil {
		*address.Flat = ""
	}
	return address, nil
}
