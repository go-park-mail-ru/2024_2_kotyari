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
	err := r.ratingUpdaterManager.UpdateProductRating(ctx, request.GetProductId())
	if err != nil {
		if errors.Is(err, errs.ProductNotFound) {
			r.log.Error("[RatingUpdaterGRPC.UpdateRating] Product not found", slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, errs.ProductNotFound.Error())
		}

		r.log.Error("[RatingUpdaterGRPC.UpdateRating] Unexpected error occurred", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, errs.InternalServerError.Error())
	}

	return &emptypb.Empty{}, nil
}
