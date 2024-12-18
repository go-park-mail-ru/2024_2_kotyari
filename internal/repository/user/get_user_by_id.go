package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (us *UsersStore) GetUserByUserID(ctx context.Context, id uint32) (model.User, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.User{}, err
	}

	us.log.Info("[UsersStore.GetUserByUserID] Started executing", slog.Any("request-id", requestID))

	const query = `
		select username, email, city, avatar_url from users where id=$1;
	`

	var user model.User

	err = us.db.QueryRow(ctx, query, id).Scan(&user.Username, &user.Email, &user.City, &user.AvatarUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			us.log.Error("[ UsersStore.GetUserByUserID ] Юзер не найден в бд", slog.String("error", err.Error()))
			return model.User{}, errs.UserDoesNotExist
		}
		us.log.Error("[ UsersStore.GetUserByUserID ] Ошибка при получении юзера", slog.String("error", err.Error()))
		return model.User{}, err
	}

	return user, nil
}
