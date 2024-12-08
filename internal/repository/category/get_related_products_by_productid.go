package category

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

func (cs *CategoriesStore) GetRelatedProductsByProductID(ctx context.Context, productID uint64, sortField string, sortOrder string) ([]model.ProductCatalog, error) {
	categories, err := cs.categoriesGetter.GetProductCategories(ctx, productID)
	if err != nil {
	}
	cs.log.Info("[CategoriesStore.GetRelatedProductsByProductID] Вывод полученных категорий", slog.Any("categories", categories))
	if len(categories) == 0 {
	}

	var allProducts []model.ProductCatalog
	seenProducts := make(map[uint32]bool)

	for _, category := range categories {

		products, err := cs.GetProductsByCategoryLink(ctx, category.LinkTo, sortField, sortOrder)
		cs.log.Info("[CategoriesStore.GetRelatedProductsByProductID] Вывод полученных продуктов n-ой категории", slog.Any("products", products))
		if err != nil {

		}

		for _, product := range products {
			if !seenProducts[product.ID] {
				allProducts = append(allProducts, product)
				seenProducts[product.ID] = true
			}
		}
		cs.log.Info("[CategoriesStore.GetRelatedProductsByProductID] Вывод просмотренных товаров", slog.Any("seenProd", seenProducts))
	}

	cs.log.Info("[CategoriesStore.GetRelatedProductsByProductID] Вывод просмотренных товаров", slog.Any("allProducts", allProducts))
	if len(allProducts) == 0 {

	}

	return allProducts, nil
}
