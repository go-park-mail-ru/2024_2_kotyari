package rorders

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (r *OrdersRepo) GetOrderByID(ctx context.Context, id uuid.UUID, userID uint32) (*model.Order, error) {
	const query = `
		SELECT o.id, o.address, o.status, o.updated_at, p.id, p.price, p.count, p.image_url, p.weight
		FROM orders o
		JOIN product_orders op ON o.id = op.order_id
		JOIN products p ON op.product_id = p.id
		WHERE o.id = $1::uuid AND o.user_id = $2;
	`

	rows, err := r.db.Query(ctx, query, id, userID)
	if err != nil {
		r.logger.Error("failed to query order by ID", slog.String("error", err.Error()), slog.String("order_id", id.String()))
		return nil, fmt.Errorf("failed to query order by ID: %w", err)
	}
	defer rows.Close()

	var ord *model.Order
	for rows.Next() {
		var (
			orderID   uuid.UUID
			address   string
			status    string
			date      time.Time
			productID string
			cost      int
			count     int
			imageURL  string
			weight    int
		)

		err := rows.Scan(&orderID, &address, &status, &date, &productID, &cost, &count, &imageURL, &weight)
		if err != nil {
			r.logger.Error("failed to scan order row", slog.String("error", err.Error()))
			return nil, fmt.Errorf("failed to scan order row: %w", err)
		}

		if ord == nil {
			ord = &model.Order{
				ID:           orderID,
				Address:      address,
				OrderStatus:  status,
				DeliveryDate: date,
				Products:     []model.OrderProduct{},
			}
		}

		ord.Products = append(ord.Products, model.OrderProduct{
			ID:       productID,
			Cost:     cost,
			Count:    count,
			ImageURL: imageURL,
			Weight:   weight,
		})
	}

	if rows.Err() != nil {
		r.logger.Error("error reading rows", slog.String("error", rows.Err().Error()))
		return nil, fmt.Errorf("error reading rows: %w", rows.Err())
	}

	if ord == nil {
		r.logger.Warn("order not found", slog.String("order_id", id.String()))
		return nil, pgx.ErrNoRows
	}

	r.logger.Info("GetOrderByID completed successfully", slog.String("order_id", id.String()))
	return ord, nil
}
