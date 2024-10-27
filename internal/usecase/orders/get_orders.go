package morders

import (
	"context"
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (m *OrdersManager) GetOrders(ctx context.Context) ([]order.Order, error) {
	var userId uint32 = 1

	return m.repo.GetOrders(ctx, userId)
}
