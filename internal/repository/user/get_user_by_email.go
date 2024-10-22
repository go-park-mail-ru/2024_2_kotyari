package user

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (ur *UsersRepo) GetUserByEmail(ctx context.Context, userModel model.User) (model.User, error) {
	const query = `
		select username, password from users where users.email =$1;	
	`

	var (
		username string
		password string
	)

	err := ur.db.QueryRow(ctx, query, userModel.Email).Scan(&username, &password)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Username: username,
		Password: password,
	}, nil
}
