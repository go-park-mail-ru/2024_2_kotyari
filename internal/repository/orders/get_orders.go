package rorders

import (
	"context"
	"fmt"
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/google/uuid"
	"log/slog"
)

func (r *OrdersRepo) GetOrders(ctx context.Context, userId uint32) ([]order.Order, error) {
	const query = `
		SELECT o.id::text, o.created_at, o.updated_at, p.id::text, p.image_url, p.short_description
		FROM orders o
		JOIN product_orders po ON o.id = po.order_id
		JOIN products p ON po.product_id = p.id
		WHERE o.user_id = $1;
	`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		r.logger.Error("failed to query orders", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to query orders: %w", err)
	}
	defer rows.Close()

	ordersMap := make(map[uuid.UUID]order.Order)

	for rows.Next() {
		var orderRow order.OrderRow

		if err = rows.Scan(&orderRow.OrderID, &orderRow.OrderDate, &orderRow.DeliveryDate, &orderRow.ProductID, &orderRow.ImageURL, &orderRow.ProductName); err != nil {
			r.logger.Error("failed to scan order row", slog.String("error", err.Error()))
			return nil, fmt.Errorf("failed to scan order row: %w", err)
		}

		if _, exists := ordersMap[orderRow.OrderID]; !exists {
			ordersMap[orderRow.OrderID] = order.Order{
				ID:           orderRow.OrderID,
				OrderDate:    orderRow.OrderDate,
				DeliveryDate: orderRow.DeliveryDate,
				Products:     []order.OrderProduct{},
			}
		}

		orders := ordersMap[orderRow.OrderID]
		orders.Products = append(orders.Products, order.OrderProduct{
			ID:       orderRow.ProductID,
			ImageURL: orderRow.ImageURL,
			Name:     orderRow.ProductName,
		})
		ordersMap[orderRow.OrderID] = orders
	}

	if rows.Err() != nil {
		r.logger.Error("error reading rows", slog.String("error", rows.Err().Error()))
		return nil, fmt.Errorf("error reading rows: %w", rows.Err())
	}

	orders := make([]order.Order, 0, len(ordersMap))
	for _, ord := range ordersMap {
		orders = append(orders, ord)
	}

	r.logger.Info("GetOrders completed successfully", slog.Int("order_count", len(orders)))
	return orders, nil
}
