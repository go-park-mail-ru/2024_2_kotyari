package profile

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/address"
	"github.com/jackc/pgx/v5"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (pr *ProfilesStore) GetProfile(ctx context.Context, id uint32) (model.Profile, error) {

	const query_profile = `
		SELECT id, 
		       email, 
		       username, 
		       age, 
		       gender, 
		       avatar_url 
		FROM users 
		WHERE users.id = $1;
	`

	var profile model.Profile

	err := pr.db.QueryRow(ctx, query_profile, id).Scan(
		&profile.ID,
		&profile.Email,
		&profile.Username,
		&profile.Age,
		&profile.Gender,
		&profile.AvatarURL,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			pr.log.Warn("[ ProfilesStore.GetProfile ] Пользователь не найден", slog.String("error", err.Error()))
			return model.Profile{}, errs.UserDoesNotExist
		}
		pr.log.Error("[ ProfilesStore.GetProfile ] Ошибка при получении профиля в бд", slog.String("error", err.Error()))
		return model.Profile{}, err
	}

	addressStore := address.AddressStore{Db: pr.db, Log: pr.log}
	address, err := addressStore.GetAddressByProfileID(ctx, profile.ID)
	if err != nil {
		if errors.Is(err, errs.AddressNotFound) {
			pr.log.Warn("[ ProfilesStore.GetProfile ] Адрес не найден для пользователя", slog.String("error", err.Error()))
			return model.Profile{}, nil
		}

		pr.log.Error("[ ProfilesStore.GetProfile ] Ошибка при получении адреса", slog.String("error", err.Error()))
		return model.Profile{}, err
	}

	if address.Flat == nil {
		*address.Flat = ""
	}

	profile.Address = address

	return profile, nil
}
