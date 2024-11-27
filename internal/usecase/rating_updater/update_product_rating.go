package rating_updater

import (
	"context"
	"errors"
	"log/slog"
	"math"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (r *RatingUpdaterService) UpdateProductRating(ctx context.Context, productID uint32) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	r.log.Info("[RatingUpdaterService.UpdateProductRating] Started executing", slog.Any("request-id", requestID))

	reviews, err := r.reviewsGetter.GetProductReviewsNoLogin(ctx, productID, utils.DefaultFieldParam, utils.DefaultOrderParam)
	if err != nil {
		if errors.Is(err, errs.NoReviewsForProduct) {
			r.log.Error("[RatingUpdaterService.UpdateProductRating] No reviews for product, setting rating to 0",
				slog.String("error", err.Error()))

			err = r.repository.UpdateProductRating(ctx, productID, 0)
			if err != nil {
				r.log.Error("[RatingUpdaterService.UpdateProductRating] Error occurred when updating product rating",
					slog.String("error", err.Error()))

				return err
			}

			return nil
		}
		r.log.Error("[RatingUpdaterService.UpdateProductRating] Error occurred when fetching products",
			slog.String("error", err.Error()))

		return err
	}

	var totalRating float32

	for _, review := range reviews.Reviews {
		totalRating += float32(review.Rating)
	}

	totalRating = totalRating / float32(len(reviews.Reviews))
	roundedTotalRating := float32(math.Round(float64(totalRating*10)) / 10)

	err = r.repository.UpdateProductRating(ctx, productID, roundedTotalRating)
	if err != nil {
		r.log.Error("[RatingUpdaterService.UpdateProductRating] Error occurred when updating product rating",
			slog.String("error", err.Error()))

		return err
	}

	return nil
}
