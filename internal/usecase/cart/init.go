package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type cartRepository interface {
	GetCartProduct(ctx context.Context, productID uint32, userID uint32) (model.CartProduct, error)
	ChangeCartProductCount(ctx context.Context, productID uint32, count int32, userID uint32) error
	RemoveCartProduct(ctx context.Context, productID uint32, count int32, userID uint32) error
	AddProduct(ctx context.Context, productID uint32, userID uint32) error
	ChangeCartProductDeletedState(ctx context.Context, productID uint32, userID uint32) error
	ChangeCartProductSelectedState(ctx context.Context, productID uint32, userID uint32, isSelected bool) error
	GetSelectedFromCart(ctx context.Context, userID uint32) (*model.CartProductsForOrderWithUser, error)
	GetSelectedCartItems(ctx context.Context, userID uint32) ([]model.ProductOrder, error)
}

type productCountGetter interface {
	GetProductCount(ctx context.Context, productID uint32) (uint32, error)
}

type CartManager struct {
	cartRepository     cartRepository
	productCountGetter productCountGetter
	log                *slog.Logger
}

func NewCartManager(repository cartRepository, productCountGetter productCountGetter, logger *slog.Logger) *CartManager {
	return &CartManager{
		cartRepository:     repository,
		productCountGetter: productCountGetter,
		log:                logger,
	}
}
