package delivery

import (
	"context"
	"errors"
	"log/slog"

	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *RatingUpdaterHandler) UpdateRating(ctx context.Context, request *ratingUpdater.UpdateRatingRequest) (*ratingUpdater.UpdateRatingResponse, error) {
	err := r.repository.UpdateProductRating(ctx, request.GetProductId(), 3)
	if err != nil {
		if errors.Is(err, errs.ProductNotFound) {
			r.log.Error("[RatingUpdaterHandler.UpdateRating] Product not found", slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, "Product not found")
		}

		r.log.Error("[RatingUpdaterHandler.UpdateRating] Unexpected error occurred", slog.String("error", err.Error()))

		return nil, status.Error(codes.NotFound, errs.InternalServerError.Error())
	}

	return &ratingUpdater.UpdateRatingResponse{
		Success: true,
		Message: "successfully changed product rating",
	}, nil
}
