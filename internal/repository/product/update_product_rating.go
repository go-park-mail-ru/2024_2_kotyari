package product

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ps *ProductsStore) UpdateProductRating(ctx context.Context, productID uint32, newRating float32) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	ps.log.Info("[ProductsStore.UpdateProductRating] Started executing", slog.Any("request-id", requestID))

	const query = `
		update products
		set rating = $2
		where id = $1;
	`

	commandTag, err := ps.db.Exec(ctx, query, productID, newRating)
	if err != nil {
		ps.log.Error("[RatingUpdaterStore.UpdateProductRating] Error occurred executing query",
			slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		ps.log.Error("[RatingUpdaterStore.UpdateProductRating] No rows were affected executing query")

		return errs.ProductNotFound
	}

	return nil
}
