package profile

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (ps *ProfilesService) GetProfile(ctx context.Context, Id uint32) (model.Profile, error) {
	ps.log.Info("Запрос на получение профиля", "userID", Id)

	profile, err := ps.profileRepo.ReadProfile(ctx, Id)
	if err != nil {
		ps.log.Error("Ошибка при получении профиля на уровне юзкейсы", "userID", Id, "error", err)
		return model.Profile{}, err
	}

	ps.log.Info("Профиль успешно получен", "userID", Id)

	return profile, nil
}
