package rorders

import (
	"context"
	"log/slog"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
)

func (r *OrdersRepo) GetOrderById(ctx context.Context, id uuid.UUID, userID uint32) (*order.Order, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	r.logger.Info("[OrdersRepo.GetOrderById] Started executing", slog.Any("request-id", requestID))

	rows, err := r.GetOrderByIdRows(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ord *order.Order
	for rows.Next() {
		var orderRow getOrderByIdRow

		err := rows.Scan(
			&orderRow.OrderID, &orderRow.Address, &orderRow.Status, &orderRow.TotalPrice, &orderRow.OrderDate,
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
				TotalPrice:   orderRow.TotalPrice,
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
