package product

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (ps *ProductsStore) GetProductsByIDs(ctx context.Context, ids []uint32) ([]model.ProductCatalog, error) {
	query := `
		SELECT id, title, price, image_url
		FROM products
		WHERE id = ANY($1) and active = TRUE and count > 0;
	`

	rows, err := ps.db.Query(ctx, query, ids)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var products []model.ProductCatalog

	for rows.Next() {
		var product model.ProductCatalog
		err = rows.Scan(
			&product.ID,
			&product.Title,
			&product.Price,
			&product.ImageURL,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return products, nil
}
