package address

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (as *AddressService) CreateUsersAddress(ctx context.Context, userID uint32, addressInfo model.Address) (uint32, error) {
	as.log.Info("Запрос на создание адреса", "userID", userID, "addressInfo", addressInfo)

	address, err := as.addressRepo.CreateAddress(ctx, userID, addressInfo)
	if err != nil {
		as.log.Error("Ошибка при создании адреса на уровне репозитория", "userID", userID, "error", err)
		return 0, fmt.Errorf("Произошла ошибка при создании адреса на уровне репозитория: %w", err)
	}

	as.log.Info("Адрес успешно создан", "userID", userID, "addressID", address)
	return address, nil
}
