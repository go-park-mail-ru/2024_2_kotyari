package user

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (us *UsersStore) CreateUser(ctx context.Context, userModel model.User) (model.User, error) {
	const insertQuery = `
		insert into users(email, username, password) 
		values ($1, $2, $3)
		returning id, city, username, avatar_url;
	`

	_, err := us.GetUserByEmail(ctx, userModel)
	if err == nil {
		us.log.Info("[ UsersStore.CreateUser ] пользователь уже существует")

		return model.User{}, errs.UserAlreadyExists
	}

	err = us.db.QueryRow(ctx, insertQuery,
		userModel.Email,
		userModel.Username,
		userModel.Password,
	).Scan(&userModel.ID, &userModel.City, &userModel.Username, &userModel.AvatarUrl)
	if err != nil {
		us.log.Error("[ UsersStore.CreateUser ] ошибка при добавлении юзера в бд", slog.String("error", err.Error()))
		return model.User{}, err
	}

	return userModel, nil
}
