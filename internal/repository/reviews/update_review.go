package reviews

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (r *ReviewsStore) UpdateReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	r.log.Info("[ReviewsStore.UpdateReview] Started executing", slog.Any("request-id", requestID))

	const query = `
		update reviews
		set text = $3, rating = $4
		where product_id = $1 and user_id = $2;
	`

	commandTag, err := r.db.Exec(ctx, query, productID, userID, review.Text, review.Rating)
	if err != nil {
		r.log.Error("[ReviewsStore.UpdateReview] Error executing query", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		r.log.Error("[ReviewsStore.UpdateReview] No rows were affected")

		return err
	}

	return nil
}
