package category

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type categoriesRepository interface {
	GetAllCategories(ctx context.Context) ([]model.Category, error)
	GetProductsByCategoryLink(ctx context.Context, categoryLink string, sortField string, sortOrder string) ([]model.ProductCatalog, error)
}

type CategoriesDelivery struct {
	repo        categoriesRepository
	log         *slog.Logger
	errResolver errs.GetErrorCode
}

func NewCategoriesDelivery(repo categoriesRepository, log *slog.Logger, errs errs.GetErrorCode) *CategoriesDelivery {
	return &CategoriesDelivery{
		repo:        repo,
		log:         log,
		errResolver: errs,
	}
}
