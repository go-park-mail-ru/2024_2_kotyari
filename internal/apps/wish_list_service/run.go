package wish_list_service

import (
	"fmt"
	"log/slog"
	"net"
)

func (app *WishlistApp) Run() error {
	lis, err := net.Listen("tcp",
		fmt.Sprintf("%s:%s", app.config.Address, app.config.Port),
	)
	if err != nil {
		app.log.Error("[  WishlistApp.Run ] ",
			slog.String("error", err.Error()),
		)

		return err
	}

	app.server.Register(app.gRPCServer)

	app.log.Info("[ WishlistApp.Run ]",
		slog.String("address", app.config.Address+app.config.Port),
	)
	if err = app.gRPCServer.Serve(lis); err != nil {
		app.log.Error("failed to serve gRPC", slog.String("error", err.Error()))
		return err
	}

	return nil
}
