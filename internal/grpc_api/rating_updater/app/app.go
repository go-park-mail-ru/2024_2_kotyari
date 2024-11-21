package app

import (
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	server *grpc.Server
	log    *slog.Logger
}

func NewApp(manager RatingUpdaterManager, logger *slog.Logger) *App {
	grpcServer := grpc.NewServer()
	Register(manager, logger, grpcServer)

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
