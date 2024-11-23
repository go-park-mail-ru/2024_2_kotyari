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

func (r *ReviewsStore) GetReview(ctx context.Context, productID uint32, userID uint32) (model.Review, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Review{}, err
	}

	r.log.Info("[ReviewsStore.GetReview] Started executing", slog.Any("request-id", requestID))

	const query = `
		select r.text, r.rating, r.is_private, r.created_at, u.username, u.avatar_url
		from reviews r
		join users u on u.id = r.user_id
		where product_id = $1 and user_id = $2;
	`

	var reviewDTO ReviewDTO

	err = r.db.QueryRow(ctx, query, productID, userID).Scan(
		&reviewDTO.Text,
		&reviewDTO.Rating,
		&reviewDTO.IsPrivate,
		&reviewDTO.CreatedAt,
		&reviewDTO.Username,
		&reviewDTO.AvatarURL,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.log.Error("[ReviewsStore.GetReview] Error no rows", slog.String("error", err.Error()))

			return model.Review{}, errs.ReviewNotFound
		}
		r.log.Error("[ReviewsStore.GetReview] Unexpected error", slog.String("error", err.Error()))

		return model.Review{}, err
	}

	return reviewDTO.ToModel(), nil
}
