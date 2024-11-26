package profile_service

import (
	"fmt"
	"log/slog"
	"net"
)

func (app *ProfilesApp) Run() error {
	lis, err := net.Listen("tcp",

		fmt.Sprintf("%s:%s", app.config.Address, app.config.Port),
	)
	if err != nil {
		app.log.Error("[  ProfilesApp.Run ] ",
			slog.String("error", err.Error()),
		)

		return err
	}

	app.delivery.Register(app.gRPCServer)

	app.log.Info("[ ProfilesApp.Run ]",
		slog.String("address", app.config.Address+app.config.Port),
	)

	if err = app.gRPCServer.Serve(lis); err != nil {
		app.log.Error("[ ProfilesApp.Run ]",
			slog.String("error", err.Error()),
		)

		return err
	}

	return nil
}
