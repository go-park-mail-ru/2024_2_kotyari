package notifications

import (
	"log/slog"
	"net/http"
)

func (n *NotificationsApp) Run() {
	n.router.HandleFunc("/ws", n.notificationsDelivery.Listen)

	slog.Info("Listening at", n.config.Port)

	http.ListenAndServe("0.0.0.0:8006", n.router)
}
