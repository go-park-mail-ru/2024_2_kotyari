package reviews

import (
	"context"
	"log/slog"

	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
)

func (r *RatingUpdaterGRPC) UpdateRating(ctx context.Context, productID uint32) error {
	_, err := r.client.UpdateRating(ctx, &ratingUpdater.UpdateRatingRequest{ProductId: productID})
	if err != nil {
		r.log.Error("[RatingUpdaterGRPC.UpdateRating] Failed to update product rating",
			slog.String("error", err.Error()))

		return err
	}

	return nil
}
