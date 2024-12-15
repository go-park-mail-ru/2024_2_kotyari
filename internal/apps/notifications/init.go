package notifications

import (
	"context"
	"log"
	"log/slog"

	notifications "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/notifications/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	notificationsGRPCLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/notifications"
	notificationsRepositoryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/notifications"
	notificationsServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/notifications"
	"google.golang.org/grpc"
)

type notificationsGRPCDelivery interface {
	Register(server *grpc.Server)
	GetOrdersUpdates(ctx context.Context, request *notifications.GetOrdersUpdatesRequest) (*notifications.GetOrdersUpdatesResponse, error)
}

type orderStatusesChanger interface {
	Run()
}

type NotificationsApp struct {
	server                    *grpc.Server
	notificationsGRPCDelivery notificationsGRPCDelivery
	orderStatusesChanger      orderStatusesChanger
	log                       *slog.Logger
	config                    configs.ServiceViperConfig
}

func NewNotificationsApp(config map[string]any) *NotificationsApp {
	cfg := configs.ParseServiceViperConfig(config)

	slogLogger := logger.InitLogger()
	db, err := postgres.LoadPgxPool()
	if err != nil {
		log.Fatal(err)
	}

	notificationsRepo := notificationsRepositoryLib.NewNotificationsStore(db, slogLogger)
	notificationsWorker := notificationsServiceLib.NewNotificationsWorker(notificationsRepo, slogLogger)
	notificationsGRPC := notificationsGRPCLib.NewNotificationsGRPC(notificationsRepo, slogLogger)

	grpcServer := grpc.NewServer()
	notificationsGRPC.Register(grpcServer)

	return &NotificationsApp{
		server:                    grpcServer,
		notificationsGRPCDelivery: notificationsGRPC,
		orderStatusesChanger:      notificationsWorker,
		log:                       slogLogger,
		config:                    cfg,
	}
}
