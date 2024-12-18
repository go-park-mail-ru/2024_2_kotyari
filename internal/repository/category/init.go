package category

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
)

type categoriesGetter interface {
	GetProductCategories(ctx context.Context, productID uint32) ([]model.Category, error)
}

type CategoriesStore struct {
	db  pool.DBPool
	log *slog.Logger
	categoriesGetter categoriesGetter
}

func NewCategoriesStore(db pool.DBPool, log *slog.Logger, categoriesGetter categoriesGetter) *CategoriesStore {
	return &CategoriesStore{
		db:               db,
		log:              log,
		categoriesGetter: categoriesGetter,
	}
}
