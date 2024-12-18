package notifications

import (
	"log/slog"
	"time"
)

const defaultOrdersUpdateTime = 2 * time.Minute

type notificationsRepo interface {
	ChangeOrdersStates() error
}

type NotificationsWorker struct {
	notificationsRepo notificationsRepo
	log               *slog.Logger
}

func NewNotificationsWorker(repo notificationsRepo, logger *slog.Logger) *NotificationsWorker {
	return &NotificationsWorker{
		notificationsRepo: repo,
		log:               logger,
	}
}
