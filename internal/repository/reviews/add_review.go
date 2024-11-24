package reviews

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (r *ReviewsStore) AddReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	r.log.Info("[ReviewsStore.AddReview] Started executing", slog.Any("request-id", requestID))

	const query = `
		insert into reviews(product_id, user_id, text, rating, is_private) 
		values ($1, $2, $3, $4, $5);
	`

	commandTag, err := r.db.Exec(ctx, query, productID, userID, review.Text, review.Rating, review.IsPrivate)

	if err != nil {
		r.log.Error("[ReviewsStore.AddReview] Error executing query", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		r.log.Error("[ReviewsStore.AddReview] No rows were affected")

		return err
	}

	return nil
}
