package profile

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func (ps *ProfilesService) UpdateProfileAvatar(ctx context.Context, id uint32, file *os.File) (string, error) {

	filepath, err := ps.imagesUsecase.SaveImage(file.Name(), file)
	if err != nil {
		ps.log.Error("[ ProfilesService.UpdateProfileAvatar ] Не удалось сохранить аватар профиля", slog.String("error", err.Error()))
		return "", fmt.Errorf("Не удалось сохранить аватар профиля: %w", err)
	}

	err = ps.profileRepo.UpdateProfileAvatar(ctx, id, filepath)
	if err != nil {
		ps.log.Error("[ ProfilesService.UpdateProfileAvatar ] Не удалось обновить профиль", slog.String("error", err.Error()))
		return "", err
	}

	return filepath, nil
}
