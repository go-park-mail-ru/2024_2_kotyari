package rorders

import (
	"context"
	"log/slog"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *OrdersRepo) scanOrderRow(rows pgx.Row) (getOrdersRow, error) {
	var orderRow getOrdersRow
	if err := rows.Scan(
		&orderRow.OrderID, &orderRow.OrderDate, &orderRow.DeliveryDate,
		&orderRow.ProductID, &orderRow.ImageURL, &orderRow.ProductName,
		&orderRow.TotalPrice, &orderRow.Status,
	); err != nil {
		r.logger.Error("[OrdersManager.GetOrders] failed to scan order row", slog.String("error", err.Error()))
		return orderRow, err
	}
	return orderRow, nil
}

func (r *OrdersRepo) updateOrdersMap(ordersMap map[uuid.UUID]map[string]order.Order, orderRow getOrdersRow) {
	deliveryKey := orderRow.DeliveryDate.Format("2024-01-01")

	if _, exists := ordersMap[orderRow.OrderID]; !exists {
		ordersMap[orderRow.OrderID] = make(map[string]order.Order)
	}

	if _, exists := ordersMap[orderRow.OrderID][deliveryKey]; !exists {
		ordersMap[orderRow.OrderID][deliveryKey] = order.Order{
			ID:           orderRow.OrderID,
			OrderDate:    orderRow.OrderDate,
			DeliveryDate: orderRow.DeliveryDate,
			TotalPrice:   orderRow.TotalPrice,
			Status:       orderRow.Status,
			Products:     []order.ProductOrder{},
		}
	}

	orderEntry := ordersMap[orderRow.OrderID][deliveryKey]
	orderEntry.Products = append(orderEntry.Products, order.ProductOrder{
		ProductID: orderRow.ProductID,
		ImageUrl:  orderRow.ImageURL,
		Name:      orderRow.ProductName,
	})

	ordersMap[orderRow.OrderID][deliveryKey] = orderEntry
}

func (r *OrdersRepo) convertOrdersMapToSlice(ordersMap map[uuid.UUID]map[string]order.Order) []order.Order {
	var orders []order.Order
	for _, orderByDeliveryDate := range ordersMap {
		for _, ord := range orderByDeliveryDate {
			orders = append(orders, ord)
		}
	}

	return orders
}

func (r *OrdersRepo) GetOrders(ctx context.Context, userID uint32) ([]order.Order, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	r.logger.Info("[OrdersRepo.GetOrders] Started executing", slog.Any("request-id", requestID))

	rows, err := r.GetOrdersRows(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[uuid.UUID]map[string]order.Order)

	for rows.Next() {
		orderRow, err := r.scanOrderRow(rows)
		if err != nil {
			return nil, err
		}

		r.updateOrdersMap(ordersMap, orderRow)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("[OrdersManager.GetOrders] error reading rows", slog.String("error", err.Error()))
		return nil, err
	}

	return r.convertOrdersMapToSlice(ordersMap), nil
}
