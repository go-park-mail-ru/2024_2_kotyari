package notifications

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	//"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (n *NotificationsStore) GetUserOrdersStates(ctx context.Context, userID uint32) ([]model.OrderState, error) {
	//requestID, err := utils.GetContextRequestID(ctx)
	//if err != nil {
	//	n.log.Error("[NotificationsStore.GetUserOrdersStates] Failed to get request-id",
	//		slog.String("error", err.Error()))
	//
	//	return nil, err
	//}
	//
	//n.log.Error("[NotificationsStore.GetUserOrdersStates] Started executing", slog.Any("request-id", requestID))

	const query = `
		select id, status
		from orders
		where user_id = $1
		and updated_at >= now() - $2::interval;
	`

	rows, err := n.db.Query(ctx, query, userID, DefaultDeliveredInterval)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			n.log.Error("[NotificationsStore.GetUserOrdersStates] No rows",
				slog.String("error", err.Error()))

			return nil, errs.NoOrders
		}

		n.log.Error("[NotificationsStore.GetUserOrdersStates] Failed to get rows",
			slog.String("error", err.Error()))

		return nil, err
	}

	orderStatuses, err := pgx.CollectRows(rows, pgx.RowToStructByName[OrderStateDTO])
	if err != nil {
		n.log.Error("[NotificationsStore.GetUserOrdersStates] Failed to convert rows to struct",
			slog.String("error", err.Error()))

		return nil, err
	}

	orderStates := make([]model.OrderState, 0, len(orderStatuses))

	for _, orderStatus := range orderStatuses {
		orderStates = append(orderStates, orderStatus.ToModel())
	}

	return orderStates, nil
}
