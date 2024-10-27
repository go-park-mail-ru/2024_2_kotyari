package orders

import (
	"context"
	"github.com/google/uuid"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model/orders"
)

type ordersManager interface {
	GetOrders(ctx context.Context) ([]order.Order, error)
	GetOrderByID(ctx context.Context, id uuid.UUID) (*order.Order, error)
}

type OrdersHandler struct {
	ordersManager ordersManager
}

func NewOrdersHandler(manager ordersManager) *OrdersHandler {
	return &OrdersHandler{ordersManager: manager}
}
