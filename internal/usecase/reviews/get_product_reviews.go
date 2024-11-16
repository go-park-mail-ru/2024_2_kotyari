package reviews

import (
	"context"
	"log/slog"
	"math"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (s *ReviewsService) GetProductReviews(ctx context.Context, productID uint32) (model.Reviews, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Reviews{}, err
	}

	s.log.Info("[ReviewsService.GetProductReviews] Started executing, requestID", slog.Any("request-id", requestID))

	reviews, err := s.reviewsRepo.GetProductReviews(ctx, productID)
	if err != nil {
		return model.Reviews{}, err
	}

	var (
		totalReviewsCount uint32
		totalRating       float32
	)

	for _, review := range reviews.Reviews {
		totalRating += float32(review.Rating)
		totalReviewsCount += 1
	}

	totalRating = totalRating / float32(len(reviews.Reviews))
	roundedTotalRating := float32(math.Round(float64(totalRating*10)) / 10)
	reviews.TotalReviewCount = totalReviewsCount
	reviews.TotalRating = roundedTotalRating

	return reviews, nil
}
