package profile

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ps *ProfilesService) ChangePassword(ctx context.Context, userId uint32, oldPassword, newPassword, repeatPassword string) error {
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
		return errs.WrongCredentials
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
