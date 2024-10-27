package product

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

const queryGetProductCategories = `
    SELECT 
        c.id, c.name,c.picture
    FROM product_categories pc
        JOIN categories c ON pc.category_id = c.id
    WHERE pc.product_id = $1 AND pc.active = true;
	`

func (ps *ProductsStore) getProductCategories(ctx context.Context, productID uint64) ([]model.Category, error) {
	rowsCategories, err := ps.db.Query(ctx, queryGetProductCategories, productID)
	if err != nil {
		ps.log.Error("[ ProductsStore.GetProductByID ] Error executing categories query", slog.String("error", err.Error()))

		return nil, err
	}
	defer rowsCategories.Close()

	var categories []model.Category

	for rowsCategories.Next() {
		var category model.Category
		err = rowsCategories.Scan(
			&category.ID, &category.Name, &category.Picture,
		)
		if err != nil {
			ps.log.Error("[ ProductsStore.GetProductByID ] Error scanning category", slog.String("error", err.Error()))
			return nil, err
		}
		categories = append(categories, category)
	}

	if rowsCategories.Err() != nil {
		ps.log.Error("[ ProductsStore.GetProductByID ] Error iterating over categories rows", slog.String("error", rowsCategories.Err().Error()))
		return nil, rowsCategories.Err()
	}

	return categories, nil
}
