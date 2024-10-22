package user

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (ur *UsersRepo) CreateUser(ctx context.Context, userModel model.User) (string, error) {
	const query = `
	insert into users(email, username, password) 
	values ($1, $2, $3)
	returning username;
	`

	var username string

	err := ur.db.QueryRow(ctx, query,
		userModel.Email,
		userModel.Username,
		userModel.Password,
	).Scan(&username)

	if err != nil {
		return "", err
	}

	return username, nil
}
