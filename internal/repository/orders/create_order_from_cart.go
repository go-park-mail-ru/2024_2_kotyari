package rorders

import (
	"context"
	"fmt"
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"

	"github.com/google/uuid"
)

const defaultStatus = "awaiting_payment"

func (r *OrdersRepo) CreateOrderFromCart(ctx context.Context, userID uint32, address string) (*order.Order, error) {
	orderID := uuid.New()

	const createOrderQuery = `
		INSERT INTO orders (id, user_id, total_price, address)
		SELECT $1, $2, COALESCE(SUM(p.price * c.count), 0), $3
		FROM carts c
		JOIN products p ON c.product_id = p.id
		WHERE c.user_id = $2
		GROUP BY c.user_id;
	`

	_, err := r.db.Exec(ctx, createOrderQuery, orderID.String(), userID, address)
	if err != nil {
		r.logger.Error("failed to create order", slog.String("error", err.Error()), slog.Uint64("user_id", uint64(userID)))
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	const insertProductsQuery = `
		INSERT INTO product_orders (order_id, product_id, option_id, count, delivery_date)
		SELECT $1, c.product_id, c.option_id, c.count, c.delivery_date
		FROM carts c
		WHERE c.user_id = $2;
	`

	_, err = r.db.Exec(ctx, insertProductsQuery, orderID.String(), userID)
	if err != nil {
		r.logger.Error("failed to insert products into order", slog.String("error", err.Error()), slog.Uint64("user_id", uint64(userID)))
		return nil, fmt.Errorf("failed to insert products into order: %w", err)
	}

	r.logger.Info("CreateOrderFromCart completed successfully", slog.String("order_id", orderID.String()), slog.Uint64("user_id", uint64(userID)))
	return &order.Order{
		ID:          orderID,
		OrderStatus: defaultStatus,
	}, nil
}
