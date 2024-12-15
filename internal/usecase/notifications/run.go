package notifications

import (
	"log/slog"
	"time"
)

func (n *NotificationsWorker) Run() {
	ticker := time.NewTicker(defaultOrdersUpdateTime)
	for {
		<-ticker.C

		if err := n.notificationsRepo.ChangeOrdersStates(); err != nil {
			n.log.Error("[NotificationsWorker.Run] Error happened changing orders states",
				slog.String("error", err.Error()))
		}
	}
}
