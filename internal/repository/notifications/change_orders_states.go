package notifications

import (
	"context"
	"log/slog"
	"time"
)

func (n *NotificationsStore) ChangeOrdersStates() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	const changeToPaidQuery = `
		update orders
		set status = $1, updated_at = now()
		where status = $2
		and created_at <= now() - $3::interval;
	`

	commandTag, err := n.db.Exec(ctx, changeToPaidQuery, Paid, AwaitingPayment, DefaultPaidInterval)
	if err != nil {
		n.log.Error("[NotificationsStore.ChangeOrdersStates] Failed to change orders statuses",
			slog.String("error", err.Error()))

		return err
	}

	n.log.Info("[NotificationsStore.ChangeOrdersStates] Number of orders changed",
		slog.Int64("orders-paid", commandTag.RowsAffected()))

	const changeToDeliveredQuery = `
		update orders
		set status = $1, updated_at = now()
		where status = $2
		and created_at <= now() - $3::interval;
	`

	commandTag, err = n.db.Exec(ctx, changeToDeliveredQuery, Delivered, Paid, DefaultDeliveredInterval)
	if err != nil {
		n.log.Error("[NotificationsStore.ChangeOrdersStates] Failed to change orders statuses",
			slog.String("error", err.Error()))

		return err
	}

	n.log.Info("[NotificationsStore.ChangeOrdersStates] Number of orders changed",
		slog.Int64("orders-delivered", commandTag.RowsAffected()))

	return nil
}
