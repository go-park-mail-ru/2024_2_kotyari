package morders

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	rorders "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/orders"
)

type OrdersRepo interface {
	GetOrders(ctx context.Context, userId uint32) (pgx.Rows, error)
	GetOrderById(ctx context.Context, id uuid.UUID, userID uint32, deliveryDate time.Time) (pgx.Rows, error)
	CreateOrderFromCart(ctx context.Context, orderData *order.OrderFromCart) (*order.Order, error)
	GetCartItems(ctx context.Context, userID uint32) ([]order.ProductOrder, error)
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

func NewOrdersManager(repo *rorders.OrdersRepo, logger *slog.Logger, cart cartGetter) *OrdersManager {
	return &OrdersManager{
		repo:   repo,
		cart:   cart,
		logger: logger,
	}
}
