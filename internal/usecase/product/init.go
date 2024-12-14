package product

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type cartManager interface {
	ProductInCart(ctx context.Context, userId uint32, productId uint32) (bool, error)
	GetCartProductCount(ctx context.Context, userID uint32, productID uint32) (uint32, error)
}

type productCardGetter interface {
	GetProductByID(ctx context.Context, productID uint32) (model.ProductCard, error)
}

type ProductService struct {
	cartManager       cartManager
	productCardGetter productCardGetter
	log               *slog.Logger
}

func NewProductService(cartManager cartManager, productCardGetter productCardGetter, logger *slog.Logger) *ProductService {
	return &ProductService{
		cartManager:       cartManager,
		productCardGetter: productCardGetter,
		log:               logger,
	}
}
