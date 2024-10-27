package morders

import (
	"context"
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (m *OrdersManager) CreateOrderFromCart(ctx context.Context, address string) (*order.Order, error) {
	var userID uint32 = 1

	return m.repo.CreateOrderFromCart(ctx, userID, address)
}

//TODO: при запросе спускать адрес для заполнения.
