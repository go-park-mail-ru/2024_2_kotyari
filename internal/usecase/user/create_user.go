package user

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

// todo сессиии тут
func (us *UsersService) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.User{}, err
	}

	us.log.Info("[UsersService.CreateUser] Started executing, requestID", slog.Any("request-id", requestID))

	user.Username = us.inputValidator.SanitizeString(user.Username)
	user.Email = us.inputValidator.SanitizeString(user.Email)
	user.Password = us.inputValidator.SanitizeString(user.Password)

	//us.log.Debug("creating new user", slog.Any("user", user))

	salt, err := utils.GenerateSalt()
	if err != nil {
		us.log.Error("[ UsersService.CreateUser ] Ошибка при генерации соли", slog.String("error", err.Error()))
		return model.User{}, err
	}

	user.Password = utils.HashPassword(user.Password, salt)
	dbUser, err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		us.log.Error("[ UsersService.CreateUser ] Ошибка при создании юзера на уровне репозитория", slog.String("error", err.Error()))
		return model.User{}, err
	}

	err = us.producer.AddPromoCode(ctx, dbUser.ID, utils.AvailPromoTenID)
	if err != nil {
		us.log.Error("[UserService.CreateUser] Failed to give user a promo",
			slog.String("error", err.Error()))
	}

	//us.log.Debug("[ UsersService.CreateUser ]", slog.String("DbUser", dbUser.Username))

	return dbUser, nil
}
