package profile

import (
	"context"
	"fmt"
	"os"
)

func (ps *ProfilesService) UpdateProfileAvatar(ctx context.Context, id uint32, file *os.File) error {

	filepath, err := ps.imagesUsecase.SaveFile(filename, file, nil)
	if err != nil {
		return fmt.Errorf("failed to save profile avatar: %w", err)
	}

	return nil
}
