package morders

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"sort"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (m *OrdersManager) processOrderRowsOfGetOrders(rows pgx.Rows) ([]order.Order, error) {
	ordersMap := make(map[uuid.UUID]map[string]order.Order)

	for rows.Next() {
		var orderRow getOrdersRow

		if err := rows.Scan(
			&orderRow.OrderID, &orderRow.OrderDate, &orderRow.DeliveryDate,
			&orderRow.ProductID, &orderRow.ImageURL, &orderRow.ProductName,
		); err != nil {
			m.logger.Error("[OrdersManager.GetOrders]  failed to scan order row", slog.String("error", err.Error()))
			return nil, err
		}

		deliveryKey := orderRow.DeliveryDate.Format("2024-01-01")

		if _, exists := ordersMap[orderRow.OrderID]; !exists {
			ordersMap[orderRow.OrderID] = make(map[string]order.Order)
		}

		if _, exists := ordersMap[orderRow.OrderID][deliveryKey]; !exists {
			ordersMap[orderRow.OrderID][deliveryKey] = order.Order{
				ID:           orderRow.OrderID,
				OrderDate:    orderRow.OrderDate,
				DeliveryDate: orderRow.DeliveryDate,
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

	if rows.Err() != nil {
		m.logger.Error("[OrdersManager.GetOrders]  error reading rows", slog.String("error", rows.Err().Error()))
		return nil, rows.Err()
	}

	var orders []order.Order
	for _, orderByDeliveryDate := range ordersMap {
		for _, ord := range orderByDeliveryDate {
			orders = append(orders, ord)
		}
	}

	sort.SliceStable(orders, func(i, j int) bool {
		if orders[i].DeliveryDate.Equal(orders[j].DeliveryDate) {
			return orders[i].OrderDate.Before(orders[j].OrderDate)
		}
		return orders[i].DeliveryDate.Before(orders[j].DeliveryDate)
	})

	return orders, nil
}

func (m *OrdersManager) GetOrders(ctx context.Context, userID uint32) ([]order.Order, error) {
	m.logger.Info("get: ", slog.Uint64("u_id", uint64(userID)))

	rows, err := m.repo.GetOrders(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders, err := m.processOrderRowsOfGetOrders(rows)
	if err != nil {
		return nil, err
	}

	m.logger.Info("GetOrders completed successfully", slog.Int("order_count", len(orders)))
	return orders, nil
}
