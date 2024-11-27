package profile

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (pr *ProfilesStore) UpdateProfileAvatar(ctx context.Context, profileID uint32, filePath string) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	pr.log.Info("[ProfilesStore.UpdateProfileAvatar] Started executing", slog.Any("request-id", requestID))

	const query = `
		UPDATE users 
		SET avatar_url = $1
		WHERE id = $2;	
	`

	_, err = pr.db.Exec(ctx, query, filePath,
		profileID)
	if err != nil {
		pr.log.Error("[ ProfilesStore.UpdateProfileAvatar ] Ошибка обновления аватара в базе данных", slog.String("error", err.Error()))
		return err
	}

	return nil
}
