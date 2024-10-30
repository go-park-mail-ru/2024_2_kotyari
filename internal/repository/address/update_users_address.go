package address

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

func (ar *AddressStore) UpdateUsersAddress(addressID uint32, addressModel model.Address) error {
	ar.log.Info("Начало обновления адреса", "addressID", addressID)

	const query = `
		update addresses set city = $1, street = $2, house = $3, flat = $4 where id = $5;
	`

	_, err := ar.db.Exec(context.Background(), query,
		addressModel.City,
		addressModel.Street,
		addressModel.House,
		addressModel.Flat,
		addressID)

	if err != nil {
		ar.log.Error("Ошибка при обновлении адреса", slog.String("error", err.Error()))
		return err
	}

	ar.log.Info("Адрес успешно обновлен", "addressID", addressID)
	return nil
}
