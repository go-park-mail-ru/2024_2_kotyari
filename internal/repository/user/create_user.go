package user

import (
	"context"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (us *UsersStore) CreateUser(ctx context.Context, userModel model.User) (model.User, error) {
	const insertQuery = `
		insert into users(email, username, password) 
		values ($1, $2, $3)
		returning id, city, avatar_url;
	`

	_, err := us.GetUserByEmail(ctx, userModel)
	if err == nil {
		return model.User{}, errs.UserAlreadyExists
	}

	err = us.db.QueryRow(ctx, insertQuery,
		userModel.Email,
		userModel.Username,
		userModel.Password,
	).Scan(&userModel.ID, &userModel.City, &userModel.AvatarUrl)
	if err != nil {
		log.Println(fmt.Errorf("[UserStore.CreateUser] An error occured: %w", err))
		return model.User{}, err
	}

	return userModel, nil
}
