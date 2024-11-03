package rorders

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

func (r *OrdersRepo) GetOrderById(ctx context.Context, id uuid.UUID, userID uint32, deliveryDate time.Time) (pgx.Rows, error) {
	const query = `
		SELECT o.id, o.address, o.status, o.created_at, u.username,
       		op.delivery_date, p.id, p.price, op.count, p.image_url, p.weight, p.title
		FROM orders o
         	JOIN users u ON o.user_id = u.id
         	JOIN product_orders op ON o.id = op.order_id
         	JOIN products p ON op.product_id = p.id
		WHERE o.id = $1::uuid 
		  	AND o.user_id = $2 
		  	AND op.delivery_date BETWEEN $3 AND $4;
	`

	startDate := deliveryDate.Truncate(time.Millisecond)
	endDate := startDate.Add(time.Millisecond)

	rows, err := r.db.Query(ctx, query, id, userID, startDate, endDate)
	if err != nil {
		r.logger.Error("[OrdersRepo.GetOrderById] failed to query order by ID", slog.String("error", err.Error()), slog.String("order_id", id.String()))
		return nil, err
	}

	return rows, nil
}
