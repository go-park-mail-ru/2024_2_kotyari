package rating_updater

import (
	"context"
	"errors"
	"log/slog"

	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (r *RatingUpdaterGRPC) UpdateRating(ctx context.Context, request *ratingUpdater.UpdateRatingRequest) (*emptypb.Empty, error) {
	err := r.manager.UpdateProductRating(ctx, request.GetProductId())
	if err != nil {
		switch {
		case errors.Is(err, errs.ProductNotFound):
			r.log.Error("[RatingUpdaterGRPC.UpdateRating] Product not found", slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, "Product not found")

		case errors.Is(err, errs.NoReviewsForProduct):
			r.log.Error("[RatingUpdaterGRPC.UpdateRating] No reviews for product", slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, "Reviews for product not found")
		default:
			r.log.Error("[RatingUpdaterGRPC.UpdateRating] Unexpected error occurred", slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, errs.InternalServerError.Error())
		}
	}

	return &emptypb.Empty{}, nil
}
