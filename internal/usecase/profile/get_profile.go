package profile

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ps *ProfilesService) GetProfile(ctx context.Context, id uint32) (model.Profile, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Profile{}, err
	}

	ps.log.Info("[ProfilesService.GetProfile] Started executing", slog.Any("request-id", requestID))

	profile, err := ps.profileRepo.GetProfile(ctx, id)
	if err != nil {
		ps.log.Error("[ ProfilesService.GetProfile ] Ошибка при получении профиля на уровне юзкейсы",
			slog.String("error", err.Error()),
		)

		return model.Profile{}, err
	}

	return profile, nil
}
