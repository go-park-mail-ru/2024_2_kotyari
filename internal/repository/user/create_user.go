package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (us *UsersStore) CreateUser(ctx context.Context, userModel model.User) (model.User, error) {
	const query = `
		insert into users(email, username, password) 
		values ($1, $2, $3)
		returning id, username, city;
	`

	var user model.User

	err := us.db.QueryRow(ctx, query,
		userModel.Email,
		userModel.Username,
		userModel.Password,
	).Scan(&user.ID, &user.Username, &user.City)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
