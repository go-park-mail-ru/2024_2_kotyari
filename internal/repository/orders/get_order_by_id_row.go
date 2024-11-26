package rorders

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *OrdersRepo) GetOrderByIdRows(ctx context.Context, id uuid.UUID, _ uint32) (pgx.Rows, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	r.logger.Info("[OrdersRepo.GetOrderByIdRows] Started executing", slog.Any("request-id", requestID))

	const query = `
		SELECT o.id, o.address, o.status, o.created_at, u.username,
       		op.delivery_date, p.id, p.price, op.count, p.image_url, p.weight, p.title
		FROM orders o
         	JOIN users u ON o.user_id = u.id
         	JOIN product_orders op ON o.id = op.order_id
         	JOIN products p ON op.product_id = p.id
		WHERE o.id = $1::uuid 
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		r.logger.Error("[OrdersRepo.GetOrderById] failed to query order by ID", slog.String("error", err.Error()), slog.String("order_id", id.String()))
		return nil, err
	}

	return rows, nil
}
