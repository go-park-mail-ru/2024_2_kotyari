package rorders

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (r *OrdersRepo) GetOrdersRows(ctx context.Context, userID uint32) (pgx.Rows, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	r.logger.Info("[OrdersRepo.GetOrdersRows] Started executing", slog.Any("request-id", requestID))

	const query = `
		SELECT o.id::uuid, o.created_at AS order_date, po.delivery_date, 
		       p.id::bigint AS product_id, p.image_url, p.title AS name, 
		       o.total_price, o.status
		FROM orders o
		JOIN product_orders po ON o.id = po.order_id
		JOIN products p ON po.product_id = p.id
		WHERE o.user_id = $1
		ORDER BY po.delivery_date DESC , o.created_at DESC ;
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		r.logger.Error("[OrdersRepo.GetOrders] failed to query orders", slog.String("error", err.Error()))
		return nil, err
	}

	return rows, nil
}
