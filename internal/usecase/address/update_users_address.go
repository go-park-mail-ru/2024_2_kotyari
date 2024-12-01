package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (as *AddressService) UpdateUsersAddress(ctx context.Context, userID uint32, newAddress model.Addresses) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	as.log.Info("[AddressService.UpdateUsersAddress] Started executing", slog.Any("request-id", requestID))

	err = utils.IsValidAddress(newAddress.Address)
	if err != nil {
		as.log.Error("[ AddressService.UpdateUsersAddress ] Неправильный формат адреса", slog.String("error", err.Error()))

		return err
	}

	confirmedAddress, err := as.addressFetcher.FetchAddress(ctx, newAddress)
	if err != nil {
		as.log.Error("[ AddressService.UpdateUsersAddress ] Ошибка при получении обновленного адреса", slog.String("error", err.Error()))

		return err
	}

	err = as.addressRepo.UpdateUsersAddress(ctx, userID, confirmedAddress)
	if err != nil {
		as.log.Error("[ AddressService.UpdateUsersAddress ] Ошибка при обновлении адреса", slog.String("error", err.Error()))
		return err
	}

	return nil
}
