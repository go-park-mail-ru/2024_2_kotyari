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
		return "", fmt.Errorf("failed to save profile avatar: %w", err)
	}

	err = ps.profileRepo.UpdateProfileAvatar(ctx, id, filepath)
	if err != nil {
		ps.log.Error("[ ProfilesService.UpdateProfile ] Не удалось обновить профиль", slog.String("error", err.Error()))
		return "", err
	}

	return filepath, nil
}
