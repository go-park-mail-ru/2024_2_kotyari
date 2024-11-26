package user

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (us *UsersService) LoginUser(ctx context.Context, user model.User) (model.User, error) {
	dbUser, err := us.userRepo.GetUserByEmail(ctx, user)
	if err != nil {
		us.log.Error("[ UsersService.LoginUser ] Не найден юзер", slog.String("error", err.Error()))
		return model.User{}, err
	}

	if !utils.VerifyPassword(dbUser.Password, user.Password) {
		us.log.Info("[ UsersService.LoginUser ] Не прошла валидация паролей")
		return model.User{}, errs.WrongCredentials
	}

	return dbUser, nil
}
