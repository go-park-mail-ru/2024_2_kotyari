package app

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/rating_updater/delivery"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	server *grpc.Server
	log    *slog.Logger
}

func NewApp(repository delivery.RatingUpdaterRepository, logger *slog.Logger, code errs.GetErrorCode) *App {
	grpcServer := grpc.NewServer()
	handler := delivery.NewRatingUpdaterHandler(repository, logger, code)
	Register(handler, grpcServer)

	return &App{
		server: grpcServer,
		log:    logger,
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", ":8004")
	if err != nil {
		return err
	}

	a.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err = a.server.Serve(l); err != nil {
		return err
	}

	return nil
}
