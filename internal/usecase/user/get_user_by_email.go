package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (us *UsersService) GetUserByEmail(ctx context.Context, user model.User) (string, string, error) {
	dbUser, err := us.userRepo.GetUserByEmail(ctx, user)
	if err != nil {
		return "", "", err
	}

	if !utils.VerifyPassword(dbUser.Password, user.Password) {
		return "", "", errs.WrongCredentials
	}

	sessionID, err := us.sessionService.Create(ctx, dbUser.ID)
	if err != nil {
		return "", "", err
	}

	return sessionID, dbUser.Username, nil
}
