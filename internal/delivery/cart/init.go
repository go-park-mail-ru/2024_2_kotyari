package cart

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type cartManager interface {
	ChangeCartProductCount(ctx context.Context, productID uint32, count int32) error
}

type cartGetter interface {
	GetCart(ctx context.Context, deliveryDate time.Time) (model.Cart, error)
}

type CartHandler struct {
	cartManager cartManager
	cartGetter  cartGetter
	log         *slog.Logger
}

func NewCartHandler(manager cartManager, getter cartGetter, logger *slog.Logger) *CartHandler {
	return &CartHandler{
		cartManager: manager,
		cartGetter:  getter,
		log:         logger,
	}
}
