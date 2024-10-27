package morders

import (
	"context"
	"github.com/google/uuid"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (m *OrdersManager) GetOrderByID(ctx context.Context, id uuid.UUID) (*order.Order, error) {
	var userID uint32 = 1

	return m.repo.GetOrderByID(ctx, id, userID)
}
