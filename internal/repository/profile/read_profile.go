package profile

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (pr *ProfilesStore) ReadProfile(ctx context.Context, Id uint32) (model.Profile, error) {
	pr.log.Info("Начало чтения профиля", "userID", Id)

	const query_profile = `
		select id, email, username, age, gender, avatar_url from users where users.id = $1;
	`

	var profile model.Profile

	err := pr.db.QueryRow(ctx, query_profile, Id).Scan(
		&profile.ID,
		&profile.Email,
		&profile.Username,
		&profile.Age,
		&profile.Gender,
		&profile.AvatarUrl,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			pr.log.Warn("Пользователь не найден", slog.String("error", err.Error()))
			return model.Profile{}, errs.UserDoesNotExist
		}
		pr.log.Error("Ошибка при чтении профиля", slog.String("error", err.Error()))
		return model.Profile{}, err
	}

	addressStore := AddressStore{db: pr.db, log: pr.log}
	address, err := addressStore.GetAddressByProfileID(ctx, profile.ID)
	if err != nil {
		if errors.Is(err, errs.AddressNotFound) {
			pr.log.Warn("Адрес не найден для пользователя", slog.String("error", err.Error()))
		} else {
			pr.log.Error("Ошибка при получении адреса", slog.String("error", err.Error()))
			return model.Profile{}, err
		}
	} else {
		profile.Address = address
	}

	pr.log.Info("Профиль успешно прочитан", "userID", Id)
	return profile, nil
}
