package profile

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (pr *ProfilesStore) GetProfile(ctx context.Context, id uint32) (model.Profile, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Profile{}, err
	}

	pr.log.Info("[ProfilesStore.GetProfile] Started executing", slog.Any("request-id", requestID))

	const queryProfile = `
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

	err = pr.db.QueryRow(ctx, queryProfile, id).Scan(
		&profile.ID,
		&profile.Email,
		&profile.Username,
		&profile.Age,
		&profile.Gender,
		&profile.AvatarURL,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			pr.log.Warn("[ ProfilesStore.GetProfile ] Пользователь не найден",
				slog.String("error", err.Error()),
			)

			return model.Profile{}, errs.UserDoesNotExist
		}

		pr.log.Error("[ ProfilesStore.GetProfile ] Ошибка при получении профиля в бд",
			slog.String("error", err.Error()),
		)

		return model.Profile{}, err
	}

	return profile, nil
}
