package rorders

import (
	"context"
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func (r *OrdersRepo) GetOrderById(ctx context.Context, id uuid.UUID, userID uint32, deliveryDate time.Time) (*order.Order, error) {
	rows, err := r.GetOrderByIdRows(ctx, id, userID, deliveryDate)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ord *order.Order
	for rows.Next() {
		var orderRow order.GetOrderByIdRow

		err := rows.Scan(
			&orderRow.OrderID, &orderRow.Address, &orderRow.Status, &orderRow.OrderDate,
			&orderRow.Username, &orderRow.Date, &orderRow.ProductID, &orderRow.Cost,
			&orderRow.Count, &orderRow.ImageURL, &orderRow.Weight, &orderRow.Title,
		)

		if err != nil {
			r.logger.Error("[OrdersManager.processOrderRows] failed to scan row", slog.String("error", err.Error()))
			return nil, err
		}

		if ord == nil {
			ord = &order.Order{
				ID:           orderRow.OrderID,
				Recipient:    orderRow.Username,
				Address:      orderRow.Address,
				Status:       orderRow.Status,
				DeliveryDate: orderRow.Date,
				OrderDate:    orderRow.OrderDate,
				Products:     []order.ProductOrder{},
			}
		}

		ord.Products = append(ord.Products, order.ProductOrder{
			ProductID: orderRow.ProductID,
			Cost:      orderRow.Cost,
			Count:     orderRow.Count,
			ImageUrl:  orderRow.ImageURL,
			Weight:    orderRow.Weight,
			Name:      orderRow.Title,
		})
	}

	if rows.Err() != nil {
		r.logger.Error("[OrdersManager.processOrderRows] error reading rows", slog.String("error", rows.Err().Error()))
		return nil, rows.Err()
	}

	return ord, nil
}
