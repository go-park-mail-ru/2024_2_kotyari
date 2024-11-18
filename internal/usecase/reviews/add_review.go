package reviews

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (s *ReviewsService) AddReview(ctx context.Context, productID uint32, userID uint32, review model.Review) (model.Review, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.Review{}, err
	}

	s.log.Info("[ReviewsService.AddReview] Started executing", slog.Any("request-id", requestID))

	if !utils.ValidateReviewRating(review) {
		return model.Review{}, errs.BadRequest
	}

	_, err = s.reviewsRepo.GetReview(ctx, productID, userID)
	if err != nil {
		if errors.Is(err, errs.ReviewNotFound) {
			s.log.Info("[ReviewsService.AddReview] Review not found, adding new one")

			review, err = s.reviewsRepo.AddReview(ctx, productID, userID, review)
			if err != nil {
				s.log.Error("[ReviewsService.AddReview] Error happened when adding review", slog.String("error", err.Error()))

				return model.Review{}, err
			}

			return review, nil
		}

		s.log.Error("[ReviewsService.AddReview] Unexpected Error happened when fetching review", slog.String("error", err.Error()))

		return model.Review{}, err
	}

	return model.Review{}, errs.ReviewAlreadyExists
}
