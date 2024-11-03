package profile

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/image"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type profileRepository interface {
	GetProfile(ctx context.Context, Id uint32) (model.Profile, error)
	UpdateProfile(ctx context.Context, profileID uint32, profileModel model.Profile) error
}

type ProfilesService struct {
	imagesUsecase *image.ImagesUsecase
	profileRepo   profileRepository
	log           *slog.Logger
}

func NewProfileService(imagesUsecase *image.ImagesUsecase, profileRepository profileRepository, logger *slog.Logger) *ProfilesService {
	return &ProfilesService{
		imagesUsecase: imagesUsecase,
		profileRepo:   profileRepository,
		log:           logger,
	}
}
