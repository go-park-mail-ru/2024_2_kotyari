package profile

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (ar *AddressStore) GetAddressByProfileID(ctx context.Context, profileID uint32) (model.Address, error) {
	ar.log.Info("Начало получения адреса по ID профиля", "profileID", profileID)

	const query = `
		select id, city, street, house, flat from addresses where addresses.user_id = $1;
	`

	var address model.Address
	err := ar.db.QueryRow(ctx, query, profileID).Scan(&address.Id,
		&address.City,
		&address.Street,
		&address.House,
		&address.Flat)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ar.log.Warn("Адрес не найден для данного профиля", slog.String("error", err.Error()))
			return model.Address{}, errs.AddressNotFound
		}
		ar.log.Error("Ошибка при получении адреса", slog.String("error", err.Error()))
		return model.Address{}, err
	}

	ar.log.Info("Адрес успешно получен", "profileID", profileID)
	return address, nil
}
