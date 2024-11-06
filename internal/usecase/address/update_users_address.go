package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (as *AddressService) UpdateUsersAddress(ctx context.Context, addressID uint32, newAddress model.Address) error {

	err := as.addressRepo.UpdateUsersAddress(ctx, addressID, newAddress)
	if err != nil {
		as.log.Error("[ AddressService.UpdateUsersAddress ] Ошибка при обновлении адреса", slog.String("error", err.Error()))
		return err
	}

	return nil
}
