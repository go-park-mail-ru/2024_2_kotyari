package user

import (
	"log/slog"
	"net"
)

func (app *UsersApp) Run() error {
	lis, err := net.Listen("tcp",
		app.config.Address+app.config.Port,
	)
	if err != nil {
		app.log.Error("[  UsersApp.Run ] ",
			slog.String("error", err.Error()),
		)

		return err
	}
	//metrics, err := grpc.NewGrpcMetrics("user")
	//if err != nil {
	//	app.log.Error("Ошибка при регистрации метрики", slog.String("error", err.Error()))
	//	return err
	//}

	//metricsMiddleware := metricsMiddleware.NewGrpcMiddleware(metrics)

	app.delivery.Register(app.gRPCServer)

	app.log.Info("[ UsersApp.Run ]",
		slog.String("address", app.config.Address+app.config.Port),
	)
	if err := app.gRPCServer.Serve(lis); err != nil {
		app.log.Error("failed to serve gRPC", slog.String("error", err.Error()))
		return err
	}

	return nil
}
