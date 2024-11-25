package reviews

import (
	"context"
	"math"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (s *ReviewsService) GetProductReviewsNoLogin(ctx context.Context, productID uint32, sortField string, sortOrder string) (model.Reviews, error) {
	//requestID, err := utils.GetContextRequestID(ctx)
	//if err != nil {
	//	return model.Reviews{}, err
	//}
	//
	//s.log.Info("[ReviewsService.GetProductReviews] Started executing, requestID", slog.Any("request-id", requestID))

	reviews, err := s.reviewsRepo.GetProductReviewsNoLogin(ctx, productID, sortField, sortOrder)
	if err != nil {
		return model.Reviews{}, err
	}

	var totalRating float32

	for _, review := range reviews.Reviews {
		totalRating += float32(review.Rating)
	}

	totalRating = totalRating / float32(len(reviews.Reviews))
	roundedTotalRating := float32(math.Round(float64(totalRating*10)) / 10)
	reviews.TotalRating = roundedTotalRating

	return reviews, nil
}

func (s *ReviewsService) GetProductReviewsWithLogin(ctx context.Context, productID uint32, userID uint32, sortField string, sortOrder string) (model.Reviews, error) {
	//requestID, err := utils.GetContextRequestID(ctx)
	//if err != nil {
	//	return model.Reviews{}, err
	//}
	//
	//s.log.Info("[ReviewsService.GetProductReviews] Started executing, requestID", slog.Any("request-id", requestID))

	reviews, err := s.reviewsRepo.GetProductReviewsWithLogin(ctx, productID, userID, sortField, sortOrder)
	if err != nil {
		return model.Reviews{}, err
	}

	var totalRating float32

	if reviews.UserReview.Rating != 0 {
		totalRating += float32(reviews.UserReview.Rating)
	}

	for _, review := range reviews.Reviews {
		totalRating += float32(review.Rating)
	}

	totalRating = totalRating / float32(reviews.TotalReviewCount)
	roundedTotalRating := float32(math.Round(float64(totalRating*10)) / 10)
	reviews.TotalRating = roundedTotalRating

	return reviews, nil
}
