package profile

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

type profileRepository interface {
	GetProfile(ctx context.Context, Id uint32) (model.Profile, error)
	UpdateProfile(ctx context.Context, profileID uint32, profileModel model.Profile) error
	ChangePassword(ctx context.Context, userId uint32, newPassword string) error
}

type userGetter interface {
	GetUserByUserID(ctx context.Context, id uint32) (model.User, error)
}

type ProfilesService struct {
	profileRepo profileRepository
	userGetter  userGetter
	log         *slog.Logger
}

func NewProfileService(profileRepository profileRepository, logger *slog.Logger) *ProfilesService {
	return &ProfilesService{
		profileRepo: profileRepository,
		log:         logger,
	}
}
