package recommendations

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
	"strings"
)

func (rs *RecStore) GetRecommendations(ctx context.Context, productId uint64) ([]model.ProductCatalog, error) {
	product, err := rs.productGetter.GetProductByID(ctx, productId)
	if err != nil {
		rs.log.Error("[ RecStore.GetRecommendations ] Ошибка при получении товара, на который требуются рекомендации",
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	allProducts, err := rs.productsOfCategoryGetter.GetRelatedProductsByProductID(ctx, productId, "rating", "asc")
	if err != nil {
		rs.log.Error("[ RecStore.GetRecommendations ] Ошибка при получении товаров из тех же категорий, что и исходный продукт",
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	productWords := strings.Fields(strings.ToLower(product.Title))

	var filteredProducts []model.ProductCatalog

	for _, p := range allProducts {
		if p.ID == product.ID {
			continue
		}
		productTitleLower := strings.ToLower(p.Title)

		for _, word := range productWords {
			if strings.Contains(productTitleLower, word) {
				filteredProducts = append(filteredProducts, p)
				break
			}
		}
	}

	if len(filteredProducts) == 0 {
		return nil, fmt.Errorf("no similar products found for productId: %d", productId)
	}

	return filteredProducts, nil
}
