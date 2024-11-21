package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (us *UsersStore) GetUserByEmail(ctx context.Context, userModel model.User) (model.User, error) {
	const query = `
		select id, username, password, city from users where users.email =$1;	
	`

	var user model.User

	err := us.db.QueryRow(ctx, query, userModel.Email).
		Scan(&user.ID, &user.Username, &user.Password, &user.City)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			us.log.Error("[ UsersStore.GetUserByEmail ] Юзер не найден",
				slog.String("error", err.Error()),
			)

			return model.User{}, errs.UserDoesNotExist
		}

		us.log.Error("[ UsersStore.GetUserByEmail ] Ошибка при получении юзера из бд",
			slog.String("error", err.Error()),
		)

		return model.User{}, err
	}

	return user, nil
}
