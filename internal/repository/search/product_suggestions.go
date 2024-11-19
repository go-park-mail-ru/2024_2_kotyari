package search

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (s *SearchStore) ProductSuggestion(ctx context.Context, searchQuery string, sortField string, sortOrder string) ([]model.ProductCatalog, error) {

	fieldSortOptions := map[string]string{
		"rating": "p.rating",
		"price":  "p.price",
	}

	field, ok := fieldSortOptions[sortField]
	if !ok {
		field = "p.created_at"
	}

	sortOrder = utils.ReturnSortOrderOption(sortOrder)

	query := fmt.Sprintf(`
		SELECT p.id, p.title, p.price, p.original_price,
		       p.discount, p.image_url, p.description
		FROM products p
			where p.active = true and p.count > 0 and to_tsvector('russian', title) @@ to_tsquery('russian', $1 || ':*')
		ORDER BY %s %s;
	`, field, sortOrder)

	rows, err := s.db.Query(ctx, query, searchQuery)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.log.Info("[ProductsStore.GetAllProducts] нет продуктов в БД")
			return nil, errs.ProductsDoesNotExists
		}

		s.log.Error("[ProductsStore.GetAllProducts] Query error ", slog.String("error", err.Error()))
	}

	defer rows.Close()

	var products []model.ProductCatalog
	for rows.Next() {
		var p model.ProductCatalog

		err = rows.Scan(
			&p.ID, &p.Title, &p.Price, &p.OriginalPrice,
			&p.Discount, &p.ImageURL, &p.Description,
		)
		if err != nil {
			s.log.Error("[ProductsStore.GetAllProducts] rows.Scan error ", slog.String("error", err.Error()))

			return nil, err
		}

		products = append(products, p)
	}

	s.log.Info("[ProductsStore.GetAllProducts] success get products", slog.Int("count", len(products)))

	return products, nil
}
