package notifications

import (
	"log/slog"

	notifications "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/notifications/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

type NotificationsHandler struct {
	client      notifications.NotificationsClient
	errResolver errs.GetErrorCode
	log         *slog.Logger
}

func NewNotificationsDelivery(client notifications.NotificationsClient, errResolver errs.GetErrorCode, logger *slog.Logger) *NotificationsHandler {
	return &NotificationsHandler{
		client:      client,
		errResolver: errResolver,
		log:         logger,
	}
}
