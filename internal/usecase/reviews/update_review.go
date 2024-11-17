package reviews

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (s *ReviewsService) UpdateReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	s.log.Info("[ReviewsService.UpdateReview] Started executing", slog.Any("request-id", requestID))

	if !utils.ValidateReviewRating(review) {
		return errs.BadRequest
	}

	_, err = s.reviewsRepo.GetReview(ctx, productID, userID)
	if err != nil {
		if errors.Is(err, errs.ReviewNotFound) {
			s.log.Info("[ReviewsService.UpdateReview] Review not found")

			return err
		}

		s.log.Error("[ReviewsService.UpdateReview] Unexpected Error happened when fetching review", slog.String("error", err.Error()))

		return err
	}

	err = s.reviewsRepo.UpdateReview(ctx, productID, userID, review)
	if err != nil {
		s.log.Error("[ReviewsService.UpdateReview] Error happened when updating review", slog.String("error", err.Error()))

		return err
	}

	return nil
}
