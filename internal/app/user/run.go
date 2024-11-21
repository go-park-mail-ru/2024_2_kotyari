package user

import (
	"log/slog"
	"net"
)

func (app *UsersApp) Run(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		app.log.Error("failed to listen", slog.String("address", address), slog.String("error", err.Error()))
		return err
	}

	app.delivery.Register(app.gRPCServer)

	app.log.Info("gRPC server starting", slog.String("address", address))

	if err := app.gRPCServer.Serve(lis); err != nil {
		app.log.Error("failed to serve gRPC", slog.String("error", err.Error()))
		return err
	}

	return nil
}
