package orders

import (
	"context"
	"log/slog"
	"sort"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (m *OrdersManager) GetOrders(ctx context.Context, userID uint32) ([]order.Order, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	m.logger.Info("[OrdersManager.GetOrders] Started executing", slog.Any("request-id", requestID))

	orders, err := m.repo.GetOrders(ctx, userID)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(orders, func(i, j int) bool {
		if orders[i].DeliveryDate.Equal(orders[j].DeliveryDate) {
			return orders[i].OrderDate.Before(orders[j].OrderDate)
		}
		return orders[i].DeliveryDate.Before(orders[j].DeliveryDate)
	})

	m.logger.Info("GetOrders completed successfully", slog.Int("order_count", len(orders)))
	return orders, nil
}
