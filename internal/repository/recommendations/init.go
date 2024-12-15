package recommendations

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productGetter interface {
	GetProductByID(ctx context.Context, productID uint32) (model.ProductCard, error)
}

type productsOfCategoryGetter interface {
	GetRelatedProductsByProductID(ctx context.Context, productID uint32, sortField string, sortOrder string) ([]model.ProductCatalog, error)
}

type RecStore struct {
	db                       *pgxpool.Pool
	log                      *slog.Logger
	productGetter            productGetter
	productsOfCategoryGetter productsOfCategoryGetter
}

func NewRecRepo(db *pgxpool.Pool, logger *slog.Logger, productGetter productGetter, productsOfCategoryGetter productsOfCategoryGetter) *RecStore {
	return &RecStore{
		db:                       db,
		log:                      logger,
		productGetter:            productGetter,
		productsOfCategoryGetter: productsOfCategoryGetter,
	}
}
