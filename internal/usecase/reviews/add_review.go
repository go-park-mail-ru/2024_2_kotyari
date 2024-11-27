package reviews

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ReviewsService) AddReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	s.log.Info("[ReviewsService.AddReview] Started executing", slog.Any("request-id", requestID))

	if !utils.ValidateReviewRating(review) {
		return errs.BadRequest
	}

	_, err = s.reviewsRepo.GetReview(ctx, productID, userID)
	if err != nil {
		if errors.Is(err, errs.ReviewNotFound) {
			s.log.Info("[ReviewsService.AddReview] Review not found, adding new one")

			err = s.reviewsRepo.AddReview(ctx, productID, userID, review)
			if err != nil {
				s.log.Error("[ReviewsService.AddReview] Error happened when adding review", slog.String("error", err.Error()))

				return err
			}

			err = s.ratingUpdater.UpdateRating(ctx, productID)
			if err != nil {
				grpcErr, ok := status.FromError(err)
				if ok {
					switch grpcErr.Code() {
					case codes.NotFound:
						s.log.Error("[ReviewsService.AddReview] Product not found",
							slog.String("error", err.Error()))

						return errs.FailedToChangeProductRating
					case codes.Unavailable:
						s.log.Error("[ReviewsService.AddReview] Service unavailable",
							slog.String("error", err.Error()))

						return nil

					default:
						s.log.Error("[ReviewsService.AddReview] Error happened when updating product rating, "+
							"intentionally ignoring",
							slog.String("error", err.Error()))

						return nil
					}
				}

				s.log.Error("[ReviewsService.AddReview] Failed to retrieve error code",
					slog.String("error", err.Error()))

				return nil
			}

			return nil
		}

		s.log.Error("[ReviewsService.AddReview] Unexpected Error happened when getting review", slog.String("error", err.Error()))

		return err
	}

	return errs.ReviewAlreadyExists
}
