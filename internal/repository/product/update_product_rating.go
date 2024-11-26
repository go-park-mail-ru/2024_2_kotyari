package product

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

func (r *ProductsStore) UpdateProductRating(ctx context.Context, productID uint32, newRating float32) error {
	const query = `
		update products
		set rating = $2
		where id = $1;
	`

	commandTag, err := r.db.Exec(ctx, query, productID, newRating)
	if err != nil {
		r.log.Error("[RatingUpdaterStore.UpdateProductRating] Error occurred executing query",
			slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		r.log.Error("[RatingUpdaterStore.UpdateProductRating] No rows were affected executing query")

		return errs.ProductNotFound
	}

	return nil
}
