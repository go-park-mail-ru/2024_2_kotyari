package category

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (cs *CategoriesStore) GetProductsByCategoryLink(ctx context.Context, categoryLink string) ([]model.ProductCatalog, error) {
	const query = `SELECT p.id, p.title, p.price, p.original_price,
						   p.discount, p.image_url, p.description
					FROM products p
						JOIN product_categories pc ON p.id = pc.product_id
						JOIN categories c ON pc.category_id = c.id
					WHERE p.active = true AND p.count > 0 AND c.link_to = $1  -- Parameter for category name
					ORDER BY p.created_at DESC;
					`

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
