package reviews

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (r *ReviewsStore) DeleteReview(ctx context.Context, productID uint32, userID uint32) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	r.log.Info("[ReviewsStore.DeleteReview] Started executing", slog.Any("request-id", requestID))

	const query = `
		delete from reviews
		where product_id = $1 and user_id = $2;
	`

	commandTag, err := r.db.Exec(ctx, query, productID, userID)
	if err != nil {
		r.log.Error("[ReviewsStore.DeleteReview] Error executing query", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		r.log.Error("[ReviewsStore.DeleteReview] No rows were affected")

		return err
	}

	return nil
}
