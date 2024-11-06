package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (us *UsersService) LoginUser(ctx context.Context, user model.User) (string, model.User, error) {
	dbUser, err := us.userRepo.GetUserByEmail(ctx, user)
	if err != nil {
		return "", model.User{}, err
	}

	if !utils.VerifyPassword(dbUser.Password, user.Password) {
		return "", model.User{}, errs.WrongCredentials
	}

	sessionID, err := us.sessionCreator.Create(ctx, dbUser.ID)
	if err != nil {
		return "", model.User{}, err
	}

	return sessionID, dbUser, nil
}
