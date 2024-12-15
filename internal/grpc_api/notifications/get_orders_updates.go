package notifications

import (
	"context"
	"errors"
	"log/slog"

	notifications "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/notifications/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *NotificationsGRPC) GetOrdersUpdates(ctx context.Context, request *notifications.GetOrdersUpdatesRequest) (*notifications.GetOrdersUpdatesResponse, error) {
	//requestID, err := utils.GetContextRequestID(ctx)
	//if err != nil {
	//	n.log.Error("[NotificationsStore.GetUserOrdersStates] Failed to get request-id",
	//		slog.String("error", err.Error()))
	//
	//	return nil, err
	//}
	//
	//n.log.Error("[NotificationsStore.GetUserOrdersStates] Started executing", slog.Any("request-id", requestID))

	states, err := n.userOrdersStatesGetter.GetUserOrdersStates(ctx, request.GetUserId())
	if err != nil {
		if errors.Is(err, errs.NoOrdersUpdates) {
			n.log.Error("[NotificationsGRPC.GetOrdersUpdates] No orders for user")

			return nil, status.Error(codes.NotFound, err.Error())
		}

		n.log.Error("[NotificationsGRPC.GetOrdersUpdates] Unexpected error", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, err.Error())
	}

	ordersUpdates := make([]*notifications.OrderUpdateMessage, 0, len(states))

	for _, state := range states {
		ordersUpdates = append(ordersUpdates, orderUpdateToGRPC(state))
	}

	return &notifications.GetOrdersUpdatesResponse{
		OrdersUpdates: ordersUpdates,
	}, nil
}
