package notifications

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (n *NotificationsStore) GetUserOrdersStates(ctx context.Context, userID uint32) ([]model.OrderState, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		n.log.Error("[NotificationsStore.GetUserOrdersStates] Failed to get request-id",
			slog.String("error", err.Error()))

		return nil, err
	}

	n.log.Error("[NotificationsStore.GetUserOrdersStates] Started executing", slog.Any("request-id", requestID))

	const selectStatusesQuery = `
		select id, new_status
		from orders
		where user_id = $1
		and status != new_status;
	`

	rows, err := n.db.Query(ctx, selectStatusesQuery, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			n.log.Error("[NotificationsStore.GetUserOrdersStates] No rows",
				slog.String("error", err.Error()))

			return nil, errs.NoOrdersUpdates
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

	if len(orderStatuses) == 0 {
		return nil, errs.NoOrdersUpdates
	}

	err = n.updateStatuses(ctx, userID)
	if err != nil {
		n.log.Error("[NotificationsStore.updateStatuses] Failed to update statuses",
			slog.String("error", err.Error()))

		return nil, err
	}

	orderStates := make([]model.OrderState, 0, len(orderStatuses))

	for _, orderStatus := range orderStatuses {
		orderStates = append(orderStates, orderStatus.ToModel())
	}

	return orderStates, nil
}

func (n *NotificationsStore) updateStatuses(ctx context.Context, userID uint32) error {
	const updateStatuses = `
		update orders
		set status = new_status
		where user_id = $1;
	`

	commandTag, err := n.db.Exec(ctx, updateStatuses, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			n.log.Info("[NotificationsStore.updateStatuses] No rows to update")

			return nil
		}

		n.log.Error("[NotificationsStore.updateStatuses] Failed to update statuses",
			slog.String("error", err.Error()))

		return err
	}

	n.log.Info("[NotificationsStore.updateStatuses] Statuses affected: ", commandTag.RowsAffected())

	return nil
}
