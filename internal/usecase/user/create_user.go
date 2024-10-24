package user

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (us *UsersService) CreateUser(ctx context.Context, user model.User) (string, string, error) {
	salt, err := utils.GenerateSalt()
	if err != nil {
		return "", "", err
	}

	hashedUserPassword := utils.HashPassword(user.Password, salt)
	user.Password = hashedUserPassword
	userID, username, err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		return "", "", err
	}

	sessionID, err := us.sessionService.Create(ctx, userID)
	if err != nil {
		return "", "", err
	}

	return sessionID, username, nil
}
