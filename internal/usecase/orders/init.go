package morders

import (
	"context"
	"github.com/google/uuid"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type OrdersRepo interface {
	GetOrders(ctx context.Context, userId uint32) ([]order.Order, error)
	GetOrderByID(ctx context.Context, id uuid.UUID, userId uint32) (*order.Order, error)
	CreateOrderFromCart(ctx context.Context, id uint32, address string) (*order.Order, error)
}

type OrdersManager struct {
	repo OrdersRepo
}

// Конструктор менеджера
func NewOrdersManager(repo OrdersRepo) *OrdersManager {
	return &OrdersManager{repo: repo}
}
