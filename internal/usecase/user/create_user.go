package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (us *UsersService) CreateUser(ctx context.Context, user model.User) (string, model.User, error) {
	salt, err := utils.GenerateSalt()
	if err != nil {
		return "", model.User{}, err
	}

	hashedUserPassword := utils.HashPassword(user.Password, salt)
	user.Password = hashedUserPassword
	dbUser, err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		return "", model.User{}, err
	}

	sessionID, err := us.sessionService.Create(ctx, dbUser.ID)
	if err != nil {
		return "", model.User{}, err
	}

	return sessionID, dbUser, nil
}
