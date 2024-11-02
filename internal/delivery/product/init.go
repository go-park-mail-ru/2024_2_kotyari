package product

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type productsRepo interface {
	GetAllProducts(ctx context.Context) ([]model.ProductCatalog, error)
	GetProductCardByID(ctx context.Context, productID uint64) (model.ProductCard, error)
}

type ProductsDelivery struct {
	repo        productsRepo
	log         *slog.Logger
	errResolver errs.GetErrorCode
}

func NewProductHandler(errResolver errs.GetErrorCode, repo productsRepo, log *slog.Logger) *ProductsDelivery {
	return &ProductsDelivery{
		repo:        repo,
		log:         log,
		errResolver: errResolver,
	}
}
