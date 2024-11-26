package rorders

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (r *OrdersRepo) GetNearestDeliveryDate(ctx context.Context, userID uint32) (time.Time, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return time.Time{}, err
	}

	r.logger.Info("[OrdersRepo.GetNearestDeliveryDate] Started executing", slog.Any("request-id", requestID))

	const query = `
		SELECT MIN(po.delivery_date)
		FROM orders o
		JOIN product_orders po ON o.id = po.order_id
		WHERE o.user_id = $1 AND po.delivery_date > NOW();
	`

	var deliveryDate time.Time
	err = r.db.QueryRow(ctx, query, userID).Scan(&deliveryDate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return time.Time{}, nil
		}

		r.logger.Error("[OrdersRepo.GetNearestDeliveryDate] failed to query nearest delivery date", slog.String("error", err.Error()))
		return time.Time{}, err
	}

	return deliveryDate, nil
}
