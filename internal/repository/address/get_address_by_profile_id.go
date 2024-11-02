package address

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/jackc/pgx/v5"
)

func (ar *AddressStore) GetAddressByProfileID(ctx context.Context, profileID uint32) (model.AddressDTO, error) {

	const query = `
		SELECT id, 
		       city, 
		       street, 
		       house, 
		       flat 
		FROM addresses 
		WHERE addresses.user_id = $1;
	`

	var address model.AddressDTO
	err := ar.db.QueryRow(ctx, query, profileID).Scan(&address.Id,
		&address.City,
		&address.Street,
		&address.House,
		&address.Flat)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ar.log.Warn("[ AddressStore.GetAddressByProfileID ] Адрес не найден для данного профиля", slog.String("error", err.Error()))
			return model.AddressDTO{}, errs.AddressNotFound
		}
		ar.log.Error("[ AddressStore.GetAddressByProfileID ] Ошибка при получении адреса", slog.String("error", err.Error()))
		return model.AddressDTO{}, err
	}

	return address, nil
}
