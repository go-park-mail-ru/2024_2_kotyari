package category

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cs *CategoriesStore) GetRelatedProductsByProductID(ctx context.Context, productID uint32, sortField string, sortOrder string) ([]model.ProductCatalog, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		cs.log.Error("[CategoriesStore.GetRelatedProductsByProductID] Failed to get request id",
			slog.String("error", err.Error()))

		return nil, err
	}

	cs.log.Info("[CategoriesStore.GetRelatedProductsByProductID] Started executing",
		slog.Any("request-id", requestID))

	categories, _ := cs.categoriesGetter.GetProductCategories(ctx, productID)

	var allProducts []model.ProductCatalog
	seenProducts := make(map[uint32]bool)

	for _, category := range categories {

		products, _ := cs.GetProductsByCategoryLink(ctx, category.LinkTo, sortField, sortOrder)

		for _, product := range products {
			if !seenProducts[product.ID] {
				allProducts = append(allProducts, product)
				seenProducts[product.ID] = true
			}
		}
	}

	return allProducts, nil
}
