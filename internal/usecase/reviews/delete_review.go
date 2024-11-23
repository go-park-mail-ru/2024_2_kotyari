package reviews

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (s *ReviewsService) DeleteReview(ctx context.Context, productID uint32, userID uint32) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	s.log.Info("[ReviewsService.DeleteReview] Started executing", slog.Any("request-id", requestID))

	_, err = s.reviewsRepo.GetReview(ctx, productID, userID)
	if err != nil {
		if errors.Is(err, errs.ReviewNotFound) {
			s.log.Info("[ReviewsService.DeleteReview] Review not found")

			return err
		}

		s.log.Error("[ReviewsService.DeleteReview] Unexpected Error happened when fetching review", slog.String("error", err.Error()))

		return err
	}

	err = s.reviewsRepo.DeleteReview(ctx, productID, userID)
	if err != nil {
		s.log.Error("[ReviewsService.DeleteReview] Error happened when deleting review", slog.String("error", err.Error()))

		return err
	}

	return nil
}
