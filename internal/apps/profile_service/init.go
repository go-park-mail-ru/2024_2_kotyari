package profile_service

import (
	"context"
	"errors"
	"fmt"
	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/profile"
	grpc2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/metrics/grpc"
	metrics2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares/metrics"
	profileRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/profile"
	profileServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/profile"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log/slog"
	"net/http"
	"time"
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
	slogLog := logger.InitLogger()

	cfg, err := configs.ParseServiceViperConfig(conf)
	if err != nil {
		slogLog.Error("[NewProfilesApp] Failed to parse cfg")

		return nil, err
	}

	if cfg.Address == "" || cfg.Port == "" {
		return nil, errors.New("[ ERROR ] пустая конфигурация сервиса Profile")
	}

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		return nil, fmt.Errorf("[ ERROR ] не инициализируется бд %v", err)
	}

	errorResolver := errs.NewErrorStore()
	metrics, err := grpc2.NewGrpcMetrics("profile")
	if err != nil {
		slogLog.Error("Ошибка при регистрации метрики", slog.String("error", err.Error()))
	}

	interceptor := metrics2.NewGrpcMiddleware(*metrics, errorResolver)

	profileRepo := profileRepoLib.NewProfileRepo(dbPool, slogLog)
	profileService := profileServiceLib.NewProfileService(profileRepo, slogLog)

	delivery := profile.NewProfilesGrpc(profileRepo, profileService, slogLog)

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.ServerMetricsInterceptor))

	router := mux.NewRouter()
	router.PathPrefix("/metrics").Handler(promhttp.Handler())
	serverProm := http.Server{Handler: router, Addr: fmt.Sprintf(":%d", 8083), ReadHeaderTimeout: 10 * time.Second}

	go func() {
		if err = serverProm.ListenAndServe(); err != nil {
			slogLog.Error("fail auth.ListenAndServe")
		}
	}()

	return &ProfilesApp{
		log:        slogLog,
		gRPCServer: grpcServer,
		delivery:   delivery,
		config:     cfg,
	}, nil
}
