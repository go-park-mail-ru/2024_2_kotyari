package orders

import (
	"context"
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/google/uuid"
	"log/slog"
)

type ordersManager interface {
	GetOrders(ctx context.Context) ([]order.Order, error)
	GetOrderByID(ctx context.Context, id uuid.UUID) (*order.Order, error)
	CreateOrderFromCart(ctx context.Context, address string) (*order.Order, error)
}

type OrdersHandler struct {
	ordersManager ordersManager
	logger        *slog.Logger
}

func NewOrdersHandler(manager ordersManager, logger *slog.Logger) *OrdersHandler {
	return &OrdersHandler{
		ordersManager: manager,
		logger:        logger,
	}
}
