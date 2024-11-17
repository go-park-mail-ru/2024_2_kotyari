package category

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cs *CategoriesStore) GetProductsByCategoryLink(ctx context.Context, categoryLink string, sortField string, sortOrder string) ([]model.ProductCatalog, error) {

	fieldSortOptions := map[string]string{
		"rating": "p.rating",
		"date":   "p.created_at",
	}

	field, ok := fieldSortOptions[sortField]
	if !ok {
		field = "p.created_at"
	}

	sortOrder = utils.ReturnSortOrderOption(sortOrder)

	query := fmt.Sprintf(`SELECT p.id, p.title, p.price, p.original_price,
						   p.discount, p.image_url, p.description
					FROM products p
						JOIN product_categories pc ON p.id = pc.product_id
						JOIN categories c ON pc.category_id = c.id
					WHERE p.active = true AND p.count > 0 AND c.link_to = $1  -- Parameter for category name
					ORDER BY %s %s;
					`, field, sortOrder)

	rows, err := cs.db.Query(ctx, query, categoryLink)
	if err != nil {
		cs.log.Error("[ CategoriesStore.ProductsByCategoryLink ] ошибка выполнения запроса",
			slog.String("categoryLink", categoryLink),
			slog.String("error", err.Error()),
		)

		return nil, err
	}

	defer rows.Close()

	var products []model.ProductCatalog

	for rows.Next() {
		var p model.ProductCatalog

		err = rows.Scan(&p.ID, &p.Title, &p.Price, &p.OriginalPrice, &p.Discount, &p.ImageURL, &p.Description)
		if err != nil {
			cs.log.Error("[ CategoriesStore.GetProductsByCategoryLink ] ошибка чтения ]",
				slog.String("categoryLink", categoryLink),
				slog.String("error", err.Error()),
			)

			return nil, err
		}

		products = append(products, p)
	}

	if len(products) == 0 {
		cs.log.Info("[ CategoriesStore.ProductsByCategoryLink ] не найдены продукты",
			slog.String("categoryLink", categoryLink),
		)

		return nil, errs.ProductsDoesNotExists
	}

	return products, nil
}
