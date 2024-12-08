package orders

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type promoCodesManager interface {
	GetPromoCode(ctx context.Context, userID uint32, promoCodeName string) (model.PromoCode, error)
	DeletePromoCode(ctx context.Context, userID uint32, promoID uint32) error
}

type OrdersRepo interface {
	GetOrders(ctx context.Context, userId uint32) ([]model.Order, error)
	GetOrderById(ctx context.Context, id uuid.UUID, userID uint32) (*model.Order, error)
	CreateOrderFromCart(ctx context.Context, orderData *model.OrderFromCart) (*model.Order, error)
	GetNearestDeliveryDate(ctx context.Context, userID uint32) (time.Time, error)
}

type OrdersManager struct {
	repo              OrdersRepo
	promoCodesManager promoCodesManager
	cart              cartGetter
	logger            *slog.Logger
}

type cartGetter interface {
	GetSelectedCartItems(ctx context.Context, userID uint32) ([]model.ProductOrder, error)
}

func NewOrdersManager(repo OrdersRepo, promoCodesManager promoCodesManager, logger *slog.Logger, cart cartGetter) *OrdersManager {
	return &OrdersManager{
		repo:              repo,
		promoCodesManager: promoCodesManager,
		cart:              cart,
		logger:            logger,
	}
}
