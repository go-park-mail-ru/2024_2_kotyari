package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (us *UsersService) CreateUser(ctx context.Context, user model.User) (string, error) {
	salt, err := utils.GenerateSalt()
	if err != nil {
		return "", err
	}

	hashedUserPassword := utils.HashPassword(user.Password, salt)
	user.Password = hashedUserPassword

	return us.userRepo.CreateUser(ctx, user)
}
