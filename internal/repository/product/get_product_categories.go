package product

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

const queryGetProductCategories = `
    SELECT 
        c.id, c.name,c.picture
    FROM product_categories pc
        JOIN categories c ON pc.category_id = c.id
    WHERE pc.product_id = $1 AND pc.active = true;
	`

func (ps *ProductsStore) getProductCategories(ctx context.Context, productID uint32) ([]model.Category, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	ps.log.Info("[ProductsStore.getProductCategories] Started executing", slog.Any("request-id", requestID))

	rowsCategories, err := ps.db.Query(ctx, queryGetProductCategories, productID)
	if err != nil {
		ps.log.Error("[ ProductsStore.GetProductCardByID ] Error executing categories query", slog.String("error", err.Error()))

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
			ps.log.Error("[ ProductsStore.GetProductCardByID ] Error scanning category", slog.String("error", err.Error()))
			return nil, err
		}
		categories = append(categories, category)
	}

	if rowsCategories.Err() != nil {
		ps.log.Error("[ ProductsStore.GetProductCardByID ] Error iterating over categories rows", slog.String("error", rowsCategories.Err().Error()))
		return nil, rowsCategories.Err()
	}

	return categories, nil
}
