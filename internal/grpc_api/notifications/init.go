package notifications

import (
	"context"
	"log/slog"

	notifications "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/notifications/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type userOrdersStatesGetter interface {
	GetUserOrdersStates(ctx context.Context, userID uint32) ([]model.OrderState, error)
}

type NotificationsGRPC struct {
	notifications.UnimplementedNotificationsServer
	userOrdersStatesGetter userOrdersStatesGetter
	log                    *slog.Logger
}

func NewNotificationsGRPC(userOrdersStatesGetter userOrdersStatesGetter, logger *slog.Logger) *NotificationsGRPC {
	return &NotificationsGRPC{
		userOrdersStatesGetter: userOrdersStatesGetter,
		log:                    logger,
	}
}
