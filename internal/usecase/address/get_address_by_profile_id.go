package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (as *AddressService) GetAddressByProfileID(ctx context.Context, userID uint32) (model.Addresses, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Addresses{}, err
	}

	as.log.Info("[AddressService.GetAddressByProfileID] Started executing", slog.Any("request-id", requestID))

	address, err := as.addressRepo.GetAddressByProfileID(ctx, userID)
	if err != nil {
		as.log.Error("[ AddressService.GetAddressByProfileID ] Ошибка при получении адреса", slog.String("error", err.Error()))
		return model.Addresses{}, err
	}

	return address, nil
}
