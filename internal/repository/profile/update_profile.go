package profile

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (pr *ProfilesStore) UpdateProfile(ctx context.Context, profileID uint32, profileModel model.Profile) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	pr.log.Info("[ProfilesStore.UpdateProfile] Started executing", slog.Any("request-id", requestID))

	const query = `
		UPDATE users 
		SET email = $1, 
		    username = $2, 
		    gender = $3 
		WHERE id = $4;	
	`

	_, err = pr.db.Exec(ctx, query,
		profileModel.Email,
		profileModel.Username,
		profileModel.Gender,
		profileID,
	)
	if err != nil {
		pr.log.Error("[ ProfilesStore.UpdateProfile ] Ошибка обновления профиля в базе данных",
			slog.String("error", err.Error()),
		)

		return err
	}

	return nil
}
