package rorders

import (
	"context"
	"log/slog"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
)

const defaultStatus = "awaiting_payment"

func (r *OrdersRepo) CreateOrderFromCart(ctx context.Context, orderData *order.OrderFromCart) (*order.Order, error) {
	_, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	const createOrderQuery = `
		INSERT INTO orders (id, user_id, total_price, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING created_at;
	`

	var createdAt time.Time
	err = r.db.QueryRow(ctx, createOrderQuery, orderData.OrderID, orderData.UserID, orderData.TotalPrice, orderData.Address).Scan(&orderData.CreatedAt)
	if err != nil {
		r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to insert order", slog.String("error", err.Error()), slog.Uint64("user_id", uint64(orderData.UserID)))
		return nil, err
	}

	const insertProductQuery = `
		INSERT INTO product_orders (id, order_id, product_id, option_id, count, delivery_date)
		SELECT $1, $2, $3, $4, $5, $6
		WHERE $5 > 0;
	`

	for _, p := range orderData.Products {
		productOrderID := uuid.New()

		_, err := r.db.Exec(ctx, insertProductQuery, productOrderID, orderData.OrderID, p.ID, p.OptionID, p.Count, orderData.DeliveryDate)
		if err != nil {
			r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to insert product in order", slog.String("error", err.Error()), slog.Uint64("user_id", uint64(orderData.UserID)))
			return nil, err
		}
	}

	const removeCartItemsQuery = `
		UPDATE carts
		SET count = 0, is_deleted = true
		WHERE user_id = $1 AND is_selected = true AND is_deleted = false;
	`

	_, err = r.db.Exec(ctx, removeCartItemsQuery, orderData.UserID)
	if err != nil {
		r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to remove selected cart items", slog.String("error", err.Error()), slog.Uint64("user_id", uint64(orderData.UserID)))
		return nil, err
	}

	return &order.Order{
		ID:         orderData.OrderID,
		Address:    orderData.Address,
		Status:     defaultStatus,
		OrderDate:  createdAt,
		Products:   orderData.Products,
		TotalPrice: orderData.TotalPrice,
	}, nil
}
