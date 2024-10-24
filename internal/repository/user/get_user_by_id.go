package user

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (us *UsersStore) GetUserByUserID(ctx context.Context, id uint32) (model.User, error) {
	const query = `
		select username, city from users where id=$1;
	`

	var user model.User

	err := us.db.QueryRow(ctx, query, id).Scan(&user.Username, &user.City)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, errs.UserDoesNotExist
		}
		return model.User{}, err
	}

	return user, nil
}
