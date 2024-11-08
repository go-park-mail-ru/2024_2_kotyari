package profile

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type imageSaver interface {
	SaveImage(filename string, file *os.File) (string, error)
}

type profileRepository interface {
	GetProfile(ctx context.Context, Id uint32) (model.Profile, error)
	UpdateProfile(ctx context.Context, profileID uint32, profileModel model.Profile) error
	UpdateProfileAvatar(ctx context.Context, profileID uint32, filePath string) error
}

type ProfilesService struct {
	imagesUsecase imageSaver
	profileRepo   profileRepository
	log           *slog.Logger
}

func NewProfileService(imagesUsecase imageSaver, profileRepository profileRepository, logger *slog.Logger) *ProfilesService {
	return &ProfilesService{
		imagesUsecase: imagesUsecase,
		profileRepo:   profileRepository,
		log:           logger,
	}
}
