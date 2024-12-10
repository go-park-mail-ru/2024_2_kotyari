package orders

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type ordersManager interface {
	GetOrders(ctx context.Context, userID uint32) ([]order.Order, error)
	GetOrderById(ctx context.Context, id uuid.UUID, userID uint32) (*order.Order, error)
	CreateOrderFromCart(ctx context.Context, address string, userID uint32, promoName string) (*order.Order, error)
	GetNearestDeliveryDate(ctx context.Context, userID uint32) (time.Time, error)
}

type OrdersHandler struct {
	ordersManager ordersManager
	logger        *slog.Logger
	errResolver   errs.GetErrorCode
}

func NewOrdersHandler(manager ordersManager, logger *slog.Logger, errResolver errs.GetErrorCode) *OrdersHandler {
	return &OrdersHandler{
		ordersManager: manager,
		logger:        logger,
		errResolver:   errResolver,
	}
}
