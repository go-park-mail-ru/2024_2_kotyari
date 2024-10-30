package profile

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type profileRepository interface {
	ReadProfile(ctx context.Context, Id uint32) (model.Profile, error)
	UpdateProfile(profileID uint32, profileModel model.Profile) error
}

type ProfilesService struct {
	profileRepo profileRepository
	log         *slog.Logger
}

func NewProfileService(profileRepository profileRepository, logger *slog.Logger) *ProfilesService {
	return &ProfilesService{
		profileRepo: profileRepository,
		log:         logger,
	}
}
