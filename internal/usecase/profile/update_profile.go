package profile

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ps *ProfilesService) UpdateProfile(oldProfileData model.Profile, newProfileData model.Profile) error {
	ps.log.Info("Начало обновления профиля на слое Usecase", "userID", oldProfileData.ID)

	newProfile := oldProfileData

	if newProfileData.Email != "" {
		if !utils.IsValidEmail(newProfileData.Email) {
			ps.log.Warn("Некорректный формат email", "email", newProfileData.Email)
			return errs.InvalidEmailFormat
		}
		newProfile.Email = newProfileData.Email
		ps.log.Info("Email обновлен", "email", newProfileData.Email)
	}

	if newProfileData.Username != "" {
		if !utils.IsValidUsername(newProfileData.Username) {
			ps.log.Warn("Некорректный формат имени пользователя", "username", newProfileData.Username)
			return errs.InvalidUsernameFormat
		}
		newProfile.Username = newProfileData.Username
		ps.log.Info("Имя пользователя обновлено", "username", newProfileData.Username)
	}

	if newProfileData.Gender != "" {
		newProfile.Gender = newProfileData.Gender
		ps.log.Info("Пол обновлен", "gender", newProfileData.Gender)
	}

	err := ps.profileRepo.UpdateProfile(oldProfileData.ID, newProfile)
	if err != nil {
		ps.log.Error("Не удалось обновить профиль", "userID", oldProfileData.ID, "error", err)
		return err
	}

	ps.log.Info("Профиль успешно обновлен", "userID", oldProfileData.ID)
	return nil
}
