package profile

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

func (pr *ProfilesStore) UpdateProfile(profileID uint32, profileModel model.Profile) error {
	pr.log.Info("Начало обновления профиля", "profileID", profileID)

	const query = `
		update users set email = $1, username = $2, gender = $3 where id = $4;	
	`

	_, err := pr.db.Exec(context.Background(), query, profileModel.Email,
		profileModel.Username,
		profileModel.Gender,
		profileID)
	if err != nil {
		pr.log.Error("Ошибка обновления профиля в базе данных", slog.String("error", err.Error()))
		return err
	}

	pr.log.Info("Профиль успешно обновлен", "profileID", profileID)
	return nil
}
