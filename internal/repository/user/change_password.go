package user

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (us *UsersStore) ChangePassword(ctx context.Context, userId uint32, newPassword string) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	us.log.Info("[UsersStore.ChangePassword] Started executing", slog.Any("request-id", requestID))

	query := `
		UPDATE "users"
		SET "password" = $1,
		    "updated_at" = CURRENT_TIMESTAMP
		WHERE "id" = $2;
	`

	_, err = us.db.Exec(ctx, query, newPassword, userId)
	if err != nil {
		return err
	}

	return nil
}
