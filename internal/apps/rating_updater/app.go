package rating_updater

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"log/slog"
	"net"

	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	reviewsUpdaterDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/rating_updater/delivery"
	reviewsUpdaterServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/rating_updater/usecase"
	productRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/product"
	reviewsRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/reviews"
	"google.golang.org/grpc"
)

type RatingUpdaterDelivery interface {
	Register(server *grpc.Server)
	UpdateRating(ctx context.Context, request *ratingUpdater.UpdateRatingRequest) (*ratingUpdater.UpdateRatingResponse, error)
}

type viperConfig struct {
	address string
	port    string
}

type RatingUpdaterApp struct {
	delivery RatingUpdaterDelivery
	server   *grpc.Server
	log      *slog.Logger
	config   viperConfig
}

func NewApp(config map[string]any) (*RatingUpdaterApp, error) {
	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		slog.Error("[RatingUpdaterApp] Failed to load dbPool", err.Error())

		return nil, err
	}

	grpcServer := grpc.NewServer()
	log := logger.InitLogger()
	productsRepo := productRepoLib.NewProductsStore(dbPool, log)
	reviewsRepo := reviewsRepoLib.NewReviewsStore(dbPool, log)
	ratingUpdaterManager := reviewsUpdaterServiceLib.NewRatingUpdateService(productsRepo, reviewsRepo, log)
	ratingUpdaterDelivery := reviewsUpdaterDeliveryLib.NewRatingUpdaterGRPC(ratingUpdaterManager, log)
	ratingUpdaterDelivery.Register(grpcServer)

	cfg := viperConfig{
		address: config[configs.KeyAddress].(string),
		port:    config[configs.KeyPort].(string),
	}

	return &RatingUpdaterApp{
		delivery: ratingUpdaterDelivery,
		server:   grpcServer,
		log:      log,
		config:   cfg,
	}, nil
}

func (a *RatingUpdaterApp) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.config.address, a.config.port))
	if err != nil {
		return err
	}

	a.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err = a.server.Serve(l); err != nil {
		return err
	}

	return nil
}
