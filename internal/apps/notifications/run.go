package notifications

import (
	"fmt"
	"log/slog"
	"net/http"
)

func (n *NotificationsApp) Run() {
	n.router.HandleFunc("/ws", n.notificationsDelivery.Listen)

	slog.Info("Listening at", n.config.Port)

	http.ListenAndServe(fmt.Sprintf("%s:%s", n.config.Address, n.config.Port), n.router)
}
