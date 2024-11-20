package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

// todo сессиии тут
func (us *UsersService) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	user.Username = us.inputValidator.SanitizeString(user.Username)
	user.Email = us.inputValidator.SanitizeString(user.Email)
	user.Password = us.inputValidator.SanitizeString(user.Password)

	salt, err := utils.GenerateSalt()
	if err != nil {
		return model.User{}, err
	}

	user.Password = utils.HashPassword(user.Password, salt)
	dbUser, err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return dbUser, nil
}
