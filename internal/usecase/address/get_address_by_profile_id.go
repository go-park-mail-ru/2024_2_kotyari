package address

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (as *AddressService) GetAddressByProfileID(ctx context.Context, userID uint32) (model.Address, error) {
	as.log.Info("Запрос адреса по ID профиля", "userID", userID)

	address, err := as.addressRepo.GetAddressByProfileID(ctx, userID)
	if err != nil {
		as.log.Error("Ошибка при получении адреса", "userID", userID, "error", err)
		return model.Address{}, err
	}

	as.log.Info("Адрес успешно получен", "userID", userID)
	return address, nil
}
