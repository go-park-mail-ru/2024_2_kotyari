package rating_updater

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	grpc2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/metrics/grpc"
	metrics2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares/metrics"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"net"
	"net/http"
	"time"

	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	reviewsUpdaterDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/rating_updater"
	productRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/product"
	reviewsRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/reviews"
	reviewsUpdaterServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/rating_updater"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RatingUpdaterDelivery interface {
	Register(server *grpc.Server)
	UpdateRating(ctx context.Context, request *ratingUpdater.UpdateRatingRequest) (*emptypb.Empty, error)
}

type RatingUpdaterApp struct {
	delivery RatingUpdaterDelivery
	server   *grpc.Server
	log      *slog.Logger
	config   configs.ServiceViperConfig
}

func NewApp(config map[string]any) (*RatingUpdaterApp, error) {
	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		slog.Error("[RatingUpdaterApp] Failed to load dbPool", err.Error())

		return nil, err
	}

	log := logger.InitLogger()
	errorResolver := errs.NewErrorStore()

	metrics, err := grpc2.NewGrpcMetrics("rating_updater")
	if err != nil {
		log.Error("Ошибка при регистрации метрики", slog.String("error", err.Error()))
	}

	interceptor := metrics2.NewGrpcMiddleware(*metrics, errorResolver)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.ServerMetricsInterceptor))
	router := mux.NewRouter()
	router.PathPrefix("/metrics").Handler(promhttp.Handler())
	serverProm := http.Server{Handler: router, Addr: fmt.Sprintf(":%d", 8084), ReadHeaderTimeout: 10 * time.Second}

	go func() {
		if err := serverProm.ListenAndServe(); err != nil {
			log.Error("fail auth.ListenAndServe")
		}
	}()
	productsRepo := productRepoLib.NewProductsStore(dbPool, log)
	reviewsRepo := reviewsRepoLib.NewReviewsStore(dbPool, log)
	ratingUpdaterManager := reviewsUpdaterServiceLib.NewRatingUpdateService(productsRepo, reviewsRepo, log)
	ratingUpdaterDelivery := reviewsUpdaterDeliveryLib.NewRatingUpdaterGRPC(ratingUpdaterManager, log)
	ratingUpdaterDelivery.Register(grpcServer)

	cfg := configs.ParseServiceViperConfig(config)

	return &RatingUpdaterApp{
		delivery: ratingUpdaterDelivery,
		server:   grpcServer,
		log:      log,
		config:   cfg,
	}, nil
}

func (a *RatingUpdaterApp) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.config.Address, a.config.Port))
	if err != nil {
		return err
	}

	a.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err = a.server.Serve(l); err != nil {
		return err
	}

	return nil
}
