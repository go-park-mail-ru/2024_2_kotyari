package profile

import (
	"context"
	"log/slog"
)

func (pr *ProfilesStore) UpdateProfileAvatar(ctx context.Context, profileID uint32, filePath string) error {
	const query = `
		UPDATE users 
		SET avatar_url = $1
		WHERE id = $2;	
	`

	_, err := pr.db.Exec(ctx, query, filePath,
		profileID)
	if err != nil {
		pr.log.Error("[ ProfilesStore.UpdateProfileAvatar ] Ошибка обновления аватара в базе данных", slog.String("error", err.Error()))
		return err
	}

	return nil
}
