package notifications

import (
	"fmt"
	"log/slog"
	"net"
)

func (n *NotificationsApp) Run() error {
	go func() {
		n.log.Info("[NotificationsApp.Run] Started worker")
		n.orderStatusesChanger.Run()
	}()

	listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%s", n.config.Address, n.config.Port))
	if err != nil {
		n.log.Error("[NotificationsApp.Run] Failed to listen to tcp",
			slog.String("error", err.Error()),
		)

		return err
	}

	n.log.Info("[NotificationsApp.Run] Server started listening at: ", slog.String("port", n.config.Port))

	if err = n.server.Serve(listener); err != nil {
		n.log.Info("[NotificationsApp.Run] Failed to start listening",
			slog.String("error", err.Error()))

		return err
	}

	return nil
}
