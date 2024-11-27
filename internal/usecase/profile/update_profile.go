package profile

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ps *ProfilesService) UpdateProfile(ctx context.Context, oldProfileData model.Profile, newProfileData model.Profile) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	ps.log.Info("[ProfilesService.UpdateProfile] Started executing", slog.Any("request-id", requestID))

	newProfile := oldProfileData

	if newProfileData.Email != "" {
		if !utils.IsValidEmail(newProfileData.Email) {
			return errs.InvalidEmailFormat
		}

		newProfile.Email = newProfileData.Email
	}

	if newProfileData.Username != "" {
		if !utils.IsValidUsername(newProfileData.Username) {
			return errs.InvalidUsernameFormat
		}

		newProfile.Username = newProfileData.Username
	}

	if newProfileData.Gender != "" {
		newProfile.Gender = newProfileData.Gender
	}

	err = ps.profileRepo.UpdateProfile(ctx, oldProfileData.ID, newProfile)
	if err != nil {
		ps.log.Error("[ ProfilesService.UpdateProfile ] Не удалось обновить профиль",
			slog.String("error", err.Error()),
		)

		return err
	}

	return nil
}
