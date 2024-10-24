package user

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (ur *UsersStore) GetUserByEmail(ctx context.Context, userModel model.User) (model.User, error) {
	const query = `
		select username, password from users where users.email =$1;	
	`

	var user model.User

	err := ur.db.QueryRow(ctx, query, userModel.Email).Scan(&user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, errs.UserDoesNotExist
		}
		return model.User{}, err
	}

	return user, nil
}
