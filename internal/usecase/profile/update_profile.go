package profile

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ps *ProfilesService) UpdateProfile(ctx context.Context, oldProfileData model.Profile, newProfileData model.Profile) error {

	newProfile := oldProfileData

	if newProfileData.Email != "" {
		if !utils.IsValidEmail(newProfileData.Email) {
			ps.log.Warn("[ ProfilesService.UpdateProfile ] Некорректный формат email", "email", newProfileData.Email)
			return errs.InvalidEmailFormat
		}
		newProfile.Email = newProfileData.Email
	}

	if newProfileData.Username != "" {
		if !utils.IsValidUsername(newProfileData.Username) {
			ps.log.Warn("[ ProfilesService.UpdateProfile ] Некорректный формат имени пользователя", "username", newProfileData.Username)
			return errs.InvalidUsernameFormat
		}
		newProfile.Username = newProfileData.Username
	}

	if newProfileData.Gender != "" {
		newProfile.Gender = newProfileData.Gender
	}

	err := ps.profileRepo.UpdateProfile(ctx, oldProfileData.ID, newProfile)
	if err != nil {
		ps.log.Error("[ ProfilesService.UpdateProfile ] Не удалось обновить профиль", slog.String("error", err.Error()))
		return err
	}

	return nil
}