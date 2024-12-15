package product

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type cartChecker interface {
	ProductInCart(ctx context.Context, userId uint32, productId uint32) (bool, error)
}

type allProductsGetter interface {
	GetAllProducts(ctx context.Context) ([]model.ProductCatalog, error)
}

type productByIdGetter interface {
	GetProductByID(ctx context.Context, userID uint32, productID uint32) (model.ProductCard, error)
}

type ProductsDelivery struct {
	allProductsGetter allProductsGetter
	productByIdGetter productByIdGetter
	log               *slog.Logger
	checker           cartChecker
	errResolver       errs.GetErrorCode
}

func NewProductHandler(errResolver errs.GetErrorCode, allProductsGetter allProductsGetter,
	productByIdGetter productByIdGetter, log *slog.Logger, cartChecker cartChecker) *ProductsDelivery {
	return &ProductsDelivery{
		allProductsGetter: allProductsGetter,
		productByIdGetter: productByIdGetter,
		log:               log,
		checker:           cartChecker,
		errResolver:       errResolver,
	}
}
