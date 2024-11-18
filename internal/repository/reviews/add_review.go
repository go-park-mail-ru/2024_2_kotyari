package reviews

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (r *ReviewsStore) AddReview(ctx context.Context, productID uint32, userID uint32, review model.Review) (model.Review, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Review{}, err
	}

	r.log.Info("[ReviewsStore.AddReview] Started executing", slog.Any("request-id", requestID))

	const query = `
		with inserted_review as (
			insert into reviews(product_id, user_id, text, rating, is_private)
			values($1, $2, $3, $4, $5)
			returning user_id, created_at
		)
		select u.username, u.avatar_url, ir.created_at
		from inserted_review ir
		join users u on u.id = ir.user_id;
	`

	err = r.db.QueryRow(ctx, query, productID, userID, review.Text, review.Rating, review.IsPrivate).Scan(
		&review.Username, &review.AvatarURL, &review.CreatedAt)
	if err != nil {
		r.log.Error("[ReviewsStore.AddReview] Error executing query", slog.String("error", err.Error()))

		return model.Review{}, err
	}

	return review, nil
}
