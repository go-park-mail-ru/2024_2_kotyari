package user

import (
	"fmt"
	"log/slog"
	"net"
)

func (app *UsersApp) Run() error {
	lis, err := net.Listen("tcp",
		fmt.Sprintf("%s:%s", app.config.Address, app.config.Port),
	)
	if err != nil {
		app.log.Error("[  UsersApp.Run ] ",
			slog.String("error", err.Error()),
		)

		return err
	}

	app.delivery.Register(app.gRPCServer)

	app.log.Info("[ UsersApp.Run ]",
		slog.String("address", app.config.Address+app.config.Port),
	)
	if err = app.gRPCServer.Serve(lis); err != nil {
		app.log.Error("failed to serve gRPC", slog.String("error", err.Error()))
		return err
	}

	return nil
}
