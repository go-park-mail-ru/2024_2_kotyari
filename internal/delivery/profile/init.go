package profile

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type profileManager interface {
	GetProfile(ctx context.Context, Id uint32) (model.Profile, error)
	UpdateProfile(oldProfileData model.Profile, newProfileData model.Profile) error
}

type ProfilesDelivery struct {
	profileManager profileManager
	log            *slog.Logger
}

func NewProfilesHandler(profileManager profileManager, logger *slog.Logger) *ProfilesDelivery {
	return &ProfilesDelivery{
		profileManager: profileManager,
		log:            logger,
	}
}
