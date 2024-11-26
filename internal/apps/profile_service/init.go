package profile_service

import (
	"context"
	"errors"
	"fmt"
	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/profile"
	profileRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/profile"
	profileServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/profile"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log/slog"
)

type profilesDelivery interface {
	Register(grpcServer *grpc.Server)

	ChangePassword(ctx context.Context, in *profilegrpc.ChangePasswordRequest) (*empty.Empty, error)
	UpdateProfileAvatar(ctx context.Context, in *profilegrpc.UpdateAvatarRequest) (*empty.Empty, error)
	UpdateProfileData(ctx context.Context, in *profilegrpc.UpdateProfileDataRequest) (*empty.Empty, error)
	GetProfile(ctx context.Context, in *profilegrpc.GetProfileRequest) (*profilegrpc.GetProfileResponse, error)
}

type ProfilesApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server

	delivery profilesDelivery
	config   configs.ServiceViperConfig
}

func NewProfilesApp(
	conf map[string]any,
) (*ProfilesApp, error) {
	cfg := configs.ParseServiceViperConfig(conf)

	slogLog := logger.InitLogger()

	if cfg.Address == "" || cfg.Port == "" {
		return nil, errors.New("[ ERROR ] пустая конфигурация сервиса Profile")
	}

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		return nil, fmt.Errorf("[ ERROR ] не инициализируется бд %v", err)
	}

	profileRepo := profileRepoLib.NewProfileRepo(dbPool, slogLog)
	profileService := profileServiceLib.NewProfileService(profileRepo, slogLog)

	delivery := profile.NewProfilesGrpc(profileRepo, profileService, slogLog)

	grpcServer := grpc.NewServer()

	return &ProfilesApp{
		log:        slogLog,
		gRPCServer: grpcServer,
		delivery:   delivery,
		config:     cfg,
	}, nil
}
