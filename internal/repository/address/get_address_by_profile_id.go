package address

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (ar *AddressStore) GetAddressByProfileID(ctx context.Context, profileID uint32) (model.Address, error) {
	const query = `
		SELECT city, 
		       street, 
		       house, 
		       flat 
		FROM addresses 
		WHERE addresses.user_id = $1;
	`

	var address AddressDTO
	err := ar.Db.QueryRow(ctx, query, profileID).Scan(
		&address.City,
		&address.Street,
		&address.House,
		&address.Flat)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ar.Log.Warn("[ AddressStore.GetAddressByProfileID ] Адрес не найден для данного профиля", slog.String("error", err.Error()))

			return model.Address{}, nil
		}

		ar.Log.Error("[ AddressStore.GetAddressByProfileID ] Ошибка при получении адреса", slog.String("error", err.Error()))

		return model.Address{}, err
	}

	return address.ToModel(), nil
}
