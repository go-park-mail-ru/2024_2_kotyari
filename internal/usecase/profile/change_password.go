package profile

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ps *ProfilesService) ChangePassword(ctx context.Context, userId uint32, oldPassword, newPassword, repeatPassword string) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	ps.log.Info("[ProfilesService.ChangePassword] Started executing", slog.Any("request-id", requestID))

	user, err := ps.userRepo.GetUserByUserID(ctx, userId)
	if err != nil {
		ps.log.Error("[ ProfilesService.ChangePassword ] GetUserByUserID ]", slog.String("error", err.Error()))

		return err
	}

	if newPassword != repeatPassword {
		return errs.PasswordsDoNotMatch
	}

	user, err = ps.userRepo.GetUserByEmail(ctx, user)
	if err != nil {
		ps.log.Error("[ ProfilesService.ChangePassword ] GetUserByEmail ", slog.String("error", err.Error()))

		return err
	}

	if !utils.VerifyPassword(user.Password, oldPassword) {
		return errs.WrongPassword
	}

	salt, err := utils.GenerateSalt()
	if err != nil {
		return err
	}

	newPassword = utils.HashPassword(newPassword, salt)

	err = ps.userRepo.ChangePassword(ctx, userId, newPassword)
	if err != nil {
		return err
	}

	return nil
}
