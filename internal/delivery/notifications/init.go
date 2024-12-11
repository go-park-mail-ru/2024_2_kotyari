package notifications

import (
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

type NotificationsHandler struct {
	errResolver errs.GetErrorCode
	log         *slog.Logger
}

func NewNotificationsDelivery(errResolver errs.GetErrorCode, logger *slog.Logger) *NotificationsHandler {
	return &NotificationsHandler{
		errResolver: errResolver,
		log:         logger,
	}
}
