package profile

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type profileManager interface {
	GetProfile(ctx context.Context, id uint32) (model.Profile, error)
	UpdateProfile(ctx context.Context, oldProfileData model.Profile, newProfileData model.Profile) error
	UpdateProfileAvatar(ctx context.Context, id uint32, file *os.File) (string, error)
}

type ProfilesDelivery struct {
	profileManager profileManager
	inputValidator *utils.InputValidator
	log            *slog.Logger
}

func NewProfilesHandler(profileManager profileManager, inputValidator *utils.InputValidator, logger *slog.Logger) *ProfilesDelivery {
	return &ProfilesDelivery{
		profileManager: profileManager,
		inputValidator: inputValidator,
		log:            logger,
	}
}
