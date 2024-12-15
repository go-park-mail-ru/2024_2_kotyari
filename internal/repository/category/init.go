package category

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type categoriesGetter interface {
	GetProductCategories(ctx context.Context, productID uint32) ([]model.Category, error)
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
