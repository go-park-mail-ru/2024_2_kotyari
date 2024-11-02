package rorders

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *OrdersRepo) GetOrders(ctx context.Context, userID uint32) (pgx.Rows, error) {
	const query = `
		SELECT o.id::uuid, o.created_at AS order_date, po.delivery_date, 
		       p.id::bigint AS product_id, p.image_url, p.title AS name
		FROM orders o
		JOIN product_orders po ON o.id = po.order_id
		JOIN products p ON po.product_id = p.id
		WHERE o.user_id = $1
		ORDER BY po.delivery_date, o.created_at;
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		r.logger.Error("[OrdersRepo.GetOrders] failed to query orders", slog.String("error", err.Error()))
		return nil, fmt.Errorf("[OrdersRepo.GetOrders] failed to query orders: %w", err)
	}
	return rows, nil
}
