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

func (r *RatingUpdaterGRPC) UpdateRating(ctx context.Context, request *ratingUpdater.UpdateRatingRequest) (*ratingUpdater.UpdateRatingResponse, error) {
	err := r.manager.UpdateProductRating(ctx, request.GetProductId())
	if err != nil {
		if errors.Is(err, errs.ProductNotFound) {
			r.log.Error("[RatingUpdaterGRPC.UpdateRating] Product not found", slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, "Product not found")
		}

		if errors.Is(err, errs.NoReviewsForProduct) {
			r.log.Error("[RatingUpdaterGRPC.UpdateRating] No reviews for product", slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, "Reviews for product not found")
		}

		r.log.Error("[RatingUpdaterGRPC.UpdateRating] Unexpected error occurred", slog.String("error", err.Error()))

		return nil, status.Error(codes.NotFound, errs.InternalServerError.Error())
	}

	return &ratingUpdater.UpdateRatingResponse{
		Success: true,
	}, nil
}
