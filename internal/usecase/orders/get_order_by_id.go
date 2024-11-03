package morders

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (m *OrdersManager) processOrderRows(rows pgx.Rows) (*order.Order, error) {
	var ord *order.Order
	for rows.Next() {
		var orderRow getOrderByIdRow

		err := rows.Scan(
			&orderRow.orderID, &orderRow.address, &orderRow.status, &orderRow.orderDate,
			&orderRow.username, &orderRow.date, &orderRow.productID, &orderRow.cost,
			&orderRow.count, &orderRow.imageURL, &orderRow.weight, &orderRow.title,
		)

		if err != nil {
			m.logger.Error("[OrdersManager.processOrderRows] failed to scan row", slog.String("error", err.Error()))
			return nil, err
		}

		if ord == nil {
			ord = &order.Order{
				ID:           orderRow.orderID,
				Recipient:    orderRow.username,
				Address:      orderRow.address,
				Status:       orderRow.status,
				DeliveryDate: orderRow.date,
				OrderDate:    orderRow.orderDate,
				Products:     []order.ProductOrder{},
			}
		}

		ord.Products = append(ord.Products, order.ProductOrder{
			ProductID: orderRow.productID,
			Cost:      orderRow.cost,
			Count:     orderRow.count,
			ImageUrl:  orderRow.imageURL,
			Weight:    orderRow.weight,
			Name:      orderRow.title,
		})
	}

	if rows.Err() != nil {
		m.logger.Error("[OrdersManager.processOrderRows] error reading rows", slog.String("error", rows.Err().Error()))
		return nil, rows.Err()
	}

	return ord, nil
}

func (m *OrdersManager) GetOrderById(ctx context.Context, id uuid.UUID, deliveryDate time.Time) (*order.Order, error) {
	userID := utils.GetContextSessionUserID(ctx)

	rows, err := m.repo.GetOrderById(ctx, id, userID, deliveryDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	order, err := m.processOrderRows(rows)
	if err != nil {
		return nil, err
	}

	if order == nil {
		m.logger.Warn("[OrdersManager.GetOrderByID] order not found", slog.String("order_id", id.String()))
		return nil, pgx.ErrNoRows
	}

	m.logger.Info("[OrdersManager.GetOrderByID] GetOrderByID completed successfully", slog.String("order_id", id.String()))
	return order, nil
}
