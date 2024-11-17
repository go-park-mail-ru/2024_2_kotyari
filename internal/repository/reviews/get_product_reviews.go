package reviews

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (r *ReviewsStore) GetProductReviews(ctx context.Context, productID uint32) (model.Reviews, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Reviews{}, err
	}

	r.log.Info("[ReviewsStore.GetProductReviews] Started executing", slog.Any("request-id", requestID))

	const query = `
		select r.text, r.rating, r.is_private, u.username, u.avatar_url, r.created_at
		from reviews r
		join users u on u.id = r.user_id
		where r.product_id = $1;
	`

	rows, err := r.db.Query(ctx, query, productID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.log.Error("[ReviewsStore.GetProductReviews] No reviews")

			return model.Reviews{}, errs.NoReviewsForProduct
		}

		r.log.Info("[ReviewsStore.GetProductReviews] Unexpected error happened", slog.String("error", err.Error()))

		return model.Reviews{}, err
	}

	reviewsDTO, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ReviewDTO])

	if len(reviewsDTO) == 0 {
		return model.Reviews{}, errs.NoReviewsForProduct
	}

	if err != nil {
		r.log.Error("[ReviewsStore.GetProductReviews] Error happened when scanning rows", slog.String("error", err.Error()))

		return model.Reviews{}, err
	}

	return model.Reviews{
		Reviews: ToReviewModelSlice(reviewsDTO),
	}, nil
}
