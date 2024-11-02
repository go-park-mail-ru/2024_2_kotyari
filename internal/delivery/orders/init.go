package orders

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type ordersManager interface {
	GetOrders(ctx context.Context) ([]order.Order, error)
	GetOrderById(ctx context.Context, id uuid.UUID, deliveryDate time.Time) (*order.Order, error)
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
