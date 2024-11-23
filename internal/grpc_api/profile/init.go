package profile

import (
	"context"
	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"google.golang.org/grpc"
	"log/slog"
)

type profileManager interface {
	GetProfile(ctx context.Context, id uint32) (model.Profile, error)
	UpdateProfile(ctx context.Context, oldProfileData model.Profile, newProfileData model.Profile) error
	ChangePassword(ctx context.Context, userId uint32, oldPassword, newPassword, repeatPassword string) error
}

type avatarSaver interface {
	UpdateProfileAvatar(ctx context.Context, profileID uint32, filePath string) error
}

type ProfilesGrpc struct {
	profilegrpc.UnimplementedProfileServer

	log         *slog.Logger
	manager     profileManager
	avatarSaver avatarSaver
}

func (p *ProfilesGrpc) Register(grpcServer *grpc.Server) {
	profilegrpc.RegisterProfileServer(grpcServer, p)
}

func NewProfilesGrpc(
	avatarSaver avatarSaver,
	manager profileManager,
	log *slog.Logger,
) *ProfilesGrpc {
	return &ProfilesGrpc{
		avatarSaver: avatarSaver,
		manager:     manager,
		log:         log,
	}
}
