package user

import (
	"context"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (us *UsersStore) CreateUser(ctx context.Context, userModel model.User) (model.User, error) {
	const selectQuery = `
		select id, city from users where email=$1;
	`

	const insertQuery = `
		insert into users(email, username, password) 
		values ($1, $2, $3)
		returning id, username, city;
	`

	var user model.User

	err := us.db.QueryRow(ctx, selectQuery, userModel.Email).Scan(&user.ID, &user.City)
	if err == nil {
		return model.User{}, errs.UserAlreadyExists
	}

	err = us.db.QueryRow(ctx, insertQuery,
		userModel.Email,
		userModel.Username,
		userModel.Password,
	).Scan(&user.ID, &user.Username, &user.City)

	if err != nil {
		log.Println(fmt.Errorf("[UserStore.CreateUser] An error occured: %w", err))
		return model.User{}, err
	}

	return user, nil
}
