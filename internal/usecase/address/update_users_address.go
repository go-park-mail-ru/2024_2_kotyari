package address

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (as *AddressService) UpdateUsersAddress(addressID uint32, newAddress model.Address) error {
	as.log.Info("Начало обновления адреса", "addressID", addressID, "newAddress", newAddress)

	err := as.addressRepo.UpdateUsersAddress(addressID, newAddress)
	if err != nil {
		as.log.Error("Ошибка при обновлении адреса", "addressID", addressID, "error", err)
		return err
	}

	as.log.Info("Адрес успешно обновлён", "addressID", addressID)
	return nil
}
