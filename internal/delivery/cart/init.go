package cart

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type cartManager interface {
	ChangeCartProductCount(ctx context.Context, productID uint32, count int32, userID uint32) error
	AddProduct(ctx context.Context, productID uint32, userID uint32) error
	RemoveProduct(ctx context.Context, productID uint32, userID uint32) error
}

type cartGetter interface {
	GetCart(ctx context.Context, deliveryDate time.Time) (model.Cart, error)
}

type CartHandler struct {
	cartManager     cartManager
	cartManipulator cartGetter
	errResolver     errs.GetErrorCode
	log             *slog.Logger
}

func NewCartHandler(manager cartManager, getter cartGetter, errorHandler errs.GetErrorCode, logger *slog.Logger) *CartHandler {
	return &CartHandler{
		cartManager:     manager,
		cartManipulator: getter,
		errResolver:     errorHandler,
		log:             logger,
	}
}
