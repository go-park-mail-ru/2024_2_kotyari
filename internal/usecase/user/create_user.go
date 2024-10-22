package user

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (us *UsersService) CreateUser(ctx context.Context, user model.User) (string, error) {
	salt, err := utils.GenerateSalt()
	if err != nil {
		return "", err
	}

	hashedUserPassword := utils.HashPassword(user.Password, salt)
	user.Password = hashedUserPassword

	return us.userManager.CreateUser(ctx, user)
}
