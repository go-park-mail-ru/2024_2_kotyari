package user

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (us *UsersService) GetUserByEmail(ctx context.Context, user model.User) (string, error) {
	dbUser, err := us.userManager.GetUserByEmail(ctx, user)
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(dbUser.Password, user.Password) {
		return "", errs.WrongCredentials
	}

	return dbUser.Username, nil
}
