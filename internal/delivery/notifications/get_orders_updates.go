package notifications

import (
	"log/slog"
	"net/http"

	notifications "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/notifications/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *NotificationsHandler) GetOrdersUpdates(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		n.log.Error("[NotificationsHandler.GetOrdersUpdates] No request ID")
		utils.WriteErrorJSONByError(w, err, n.errResolver)

		return
	}

	n.log.Info("[NotificationsHandler.GetOrdersUpdates] AddReview handler",
		slog.Any("request-id", requestID))

	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		n.log.Error("[NotificationsHandler.GetOrdersUpdates] No UserID")
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	newCTX, err := utils.AddMetadataRequestID(r.Context())
	if err != nil {
		n.log.Error("[NotificationsHandler.GetOrdersUpdates] Failed to imbue context with request-id",
			slog.String("error", err.Error()))
		utils.WriteErrorJSONByError(w, err, n.errResolver)

		return
	}

	ordersUpdates, err := n.client.GetOrdersUpdates(newCTX, &notifications.GetOrdersUpdatesRequest{UserId: userID})

	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				n.log.Error("[NotificationsHandler.GetOrdersUpdates] User has no orders updates",
					slog.String("error", err.Error()))
				utils.WriteErrorJSONByError(w, errs.NoOrdersUpdates, n.errResolver)

				return
			case codes.Unavailable:
				n.log.Error("[NotificationsHandler.GetOrdersUpdates] Service unavailable",
					slog.String("error", err.Error()))

				utils.WriteErrorJSONByError(w, errs.InternalServerError, n.errResolver)

				return

			default:
				n.log.Error("[ReviewsService.AddReview] Unexpected error",
					slog.String("error", err.Error()))
				utils.WriteErrorJSONByError(w, errs.InternalServerError, n.errResolver)

				return
			}
		}

		n.log.Error("[ReviewsService.AddReview] Failed to retrieve error code",
			slog.String("error", err.Error()))

		utils.WriteErrorJSONByError(w, errs.InternalServerError, n.errResolver)

		return
	}

	orderUpdatesSlice := make([]OrderUpdate, 0, len(ordersUpdates.OrdersUpdates))

	for _, update := range ordersUpdates.OrdersUpdates {
		orderUpdatesSlice = append(orderUpdatesSlice, orderUpdateFromGrpc(update))
	}

	utils.WriteJSON(w, http.StatusOK, OrderUpdateResponse{Updates: orderUpdatesSlice})
}
