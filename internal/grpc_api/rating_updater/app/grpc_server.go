package app

import (
	"context"
	"errors"
	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RatingUpdaterRepository interface {
	UpdateProductRating(ctx context.Context, productID uint32, newRating float32) error
}

//type RatingUpdaterDelivery interface {
//	UpdateRating(ctx context.Context, request *ratingUpdater.UpdateRatingRequest) (*ratingUpdater.UpdateRatingResponse, error)
//}

type RatingUpdaterServer struct {
	//delivery RatingUpdaterDelivery
	repository RatingUpdaterRepository
	ratingUpdater.UnimplementedRatingUpdaterServer
}

func Register(repository RatingUpdaterRepository, server *grpc.Server) {
	ratingUpdater.RegisterRatingUpdaterServer(server, &RatingUpdaterServer{repository: repository})
}

func (r *RatingUpdaterServer) UpdateRating(ctx context.Context, request *ratingUpdater.UpdateRatingRequest) (*ratingUpdater.UpdateRatingResponse, error) {

	err := r.repository.UpdateProductRating(ctx, request.GetProductId(), 3)
	if err != nil {
		if errors.Is(err, errs.ProductNotFound) {
			//r.log.Error("[RatingUpdaterHandler.UpdateRating] Product not found", slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, "Product not found")
		}

		//r.log.Error("[RatingUpdaterHandler.UpdateRating] Unexpected error occurred", slog.String("error", err.Error()))

		return nil, status.Error(codes.NotFound, errs.InternalServerError.Error())
	}

	return &ratingUpdater.UpdateRatingResponse{
		Success: true,
		Message: "successfully changed product rating",
	}, nil
}
