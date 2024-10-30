package address

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (ar *AddressStore) CreateAddress(ctx context.Context, profileID uint32, addressModel model.Address) (uint32, error) {
	ar.log.Info("Начало создания адреса", "profileID", profileID, "addressModel", addressModel)

	const query = `
		insert into address (profile_id, city, street, house, flat)
		values ($1, $2, $3, $4, $5)
		returning id;
	`

	var addressID uint32
	err := ar.db.QueryRow(ctx, query,
		profileID,
		addressModel.City,
		addressModel.Street,
		addressModel.House,
		addressModel.Flat).Scan(&addressID)

	if err != nil {
		ar.log.Error("Ошибка при создании адреса", slog.String("error", err.Error()))
		return 0, err
	}

	ar.log.Info("Адрес успешно создан", "addressID", addressID, "profileID", profileID)
	return addressID, nil
}
