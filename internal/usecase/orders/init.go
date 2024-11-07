package orders

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type OrdersRepo interface {
	GetOrders(ctx context.Context, userId uint32) ([]order.Order, error)
	GetOrderById(ctx context.Context, id uuid.UUID, userID uint32, deliveryDate time.Time) (*order.Order, error)
	CreateOrderFromCart(ctx context.Context, orderData *order.OrderFromCart) (*order.Order, error)
	GetNearestDeliveryDate(ctx context.Context, userID uint32) (time.Time, error)
}

type OrdersManager struct {
	repo   OrdersRepo
	cart   cartGetter
	logger *slog.Logger
}

type cartGetter interface {
	GetSelectedCartItems(ctx context.Context, userID uint32) ([]order.ProductOrder, error)
}

func NewOrdersManager(repo OrdersRepo, logger *slog.Logger, cart cartGetter) *OrdersManager {
	return &OrdersManager{
		repo:   repo,
		cart:   cart,
		logger: logger,
	}
}
