package recommendations

import (
	"context"
	"fmt"
	"log/slog"
	"sort"
	"strings"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (rs *RecStore) GetRecommendations(ctx context.Context, productId uint32) ([]model.ProductCatalog, error) {
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

	productTags := make(map[string]bool, len(product.Tags))
	for _, tag := range product.Tags {
		productTags[strings.ToLower(tag)] = true
	}
	rs.log.Info("[ RecStore.GetRecommendations ] теги продукта ", slog.Any("producttags", productTags))

	type scoredProduct struct {
		product model.ProductCatalog
		score   int
	}

	var scoredProducts []scoredProduct

	for _, p := range allProducts {
		if p.ID == product.ID {
			continue
		}
		rs.log.Info("[ RecStore.GetRecommendations ] type prod", slog.Any("p.type", p.Type))
		score := 0
		if strings.ToLower(p.Type) != strings.ToLower(product.Type) {
			continue
		} else {
			score += 1
		}

		for _, relatedTag := range p.Tags {
			rs.log.Info("[ RecStore.GetRecommendations ] проверка", slog.Any("проверка", productTags[strings.ToLower(relatedTag)]))
			rs.log.Info("[ RecStore.GetRecommendations ] тег релейтед", slog.Any("тег рел", relatedTag))

			if productTags[strings.ToLower(relatedTag)] {
				score++
			}
		}
		rs.log.Info("[ RecStore.GetRecommendations ] подсчет score", slog.Any("scores", scoredProducts))
		if score > 0 {
			scoredProducts = append(scoredProducts, scoredProduct{product: p, score: score})
		}
	}

	if len(scoredProducts) == 0 {
		rs.log.Warn("[ RecStore.GetRecommendations ] No similar products found",
			slog.Any("productId", productId),
		)
		return nil, fmt.Errorf("no similar products found for productId: %d", productId)
	}

	sort.Slice(scoredProducts, func(i, j int) bool {
		return scoredProducts[i].score > scoredProducts[j].score
	})

	filteredProducts := make([]model.ProductCatalog, len(scoredProducts))
	for i, sp := range scoredProducts {
		filteredProducts[i] = sp.product
	}

	return filteredProducts, nil
}
