package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (ur *UsersStore) CreateUser(ctx context.Context, userModel model.User) (uint32, string, error) {
	const query = `
		insert into users(email, username, password) 
		values ($1, $2, $3)
		returning id, username;
	`

	var (
		userID   uint32
		username string
	)

	err := ur.db.QueryRow(ctx, query,
		userModel.Email,
		userModel.Username,
		userModel.Password,
	).Scan(&userID, &username)

	if err != nil {
		return 0, "", err
	}

	return userID, username, nil
}
