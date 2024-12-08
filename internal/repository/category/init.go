package category

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type categoriesGetter interface {
	GetProductCategories(ctx context.Context, productID uint64) ([]model.Category, error)
}

type CategoriesStore struct {
	db               *pgxpool.Pool
	log              *slog.Logger
	categoriesGetter categoriesGetter
}

func NewCategoriesStore(db *pgxpool.Pool, log *slog.Logger, categoriesGetter categoriesGetter) *CategoriesStore {
	return &CategoriesStore{
		db:               db,
		log:              log,
		categoriesGetter: categoriesGetter,
	}
}
