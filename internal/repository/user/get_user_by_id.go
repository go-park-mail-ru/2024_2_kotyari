package user

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (us *UsersStore) GetUserByUserID(ctx context.Context, id uint32) (model.User, error) {
	const query = `
		select username, email, city, avatar_url from users where id=$1;
	`

	var user model.User

	err := us.db.QueryRow(ctx, query, id).Scan(&user.Username, &user.Email, &user.City, &user.AvatarUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, errs.UserDoesNotExist
		}

		log.Println(fmt.Errorf("[UserStore.GetUserByUserID] An error occured: %w", err))
		return model.User{}, err
	}

	return user, nil
}
