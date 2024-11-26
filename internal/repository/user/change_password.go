package user

import (
	"context"
)

func (us *UsersStore) ChangePassword(ctx context.Context, userId uint32, newPassword string) error {
	query := `
		UPDATE "users"
		SET "password" = $1,
		    "updated_at" = CURRENT_TIMESTAMP
		WHERE "id" = $2;
	`

	_, err := us.db.Exec(ctx, query, newPassword, userId)
	if err != nil {
		return err
	}

	return nil
}
