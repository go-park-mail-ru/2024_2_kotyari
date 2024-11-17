package profile_service

import (
	"context"
	"errors"
	"fmt"
	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/profile"
	profileRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/profile"
	profileServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/profile"
	"google.golang.org/grpc"
	"log/slog"
)

type profilesDelivery interface {
	Register(grpcServer *grpc.Server)
	GetProfile(ctx context.Context, in *profilegrpc.GetProfileRequest) (*profilegrpc.GetProfileResponse, error)
	UpdateProfileAvatar(ctx context.Context, in *profilegrpc.UpdateAvatarRequest) (*profilegrpc.UpdateAvatarResponse, error)
	UpdateProfileData(ctx context.Context, in *profilegrpc.UpdateProfileDataRequest) (*profilegrpc.UpdateProfileDataResponse, error)
}

type config struct {
	Address string
	Port    string
}

type ProfilesApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server

	delivery profilesDelivery
	config   config
}

func NewProfilesApp(
	grpcServer *grpc.Server,
	conf map[string]any,
	slogLog *slog.Logger,
) (*ProfilesApp, error) {
	c := config{
		Address: conf[configs.KeyAddress].(string),
		Port:    conf[configs.KeyPort].(string),
	}

	if c.Address == "" || c.Port == "" {
		return nil, errors.New("config is empty")
	}

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		return nil, fmt.Errorf("не инициализируется бд %v", err)
	}

	profileRepo := profileRepoLib.NewProfileRepo(dbPool, slogLog)
	profileService := profileServiceLib.NewProfileService(profileRepo, slogLog)

	delivery := profile.NewProfilesGrpc(profileRepo, profileService, slogLog)

	return &ProfilesApp{
		log:        slogLog,
		gRPCServer: grpcServer,
		delivery:   delivery,
		config:     c,
	}, nil
}
