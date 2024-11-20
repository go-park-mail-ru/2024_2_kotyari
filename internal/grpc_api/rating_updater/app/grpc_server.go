package app

import (
	"context"

	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"google.golang.org/grpc"
)

type RatingUpdaterDelivery interface {
	UpdateRating(ctx context.Context, request *ratingUpdater.UpdateRatingRequest) (*ratingUpdater.UpdateRatingResponse, error)
}

type RatingUpdaterServer struct {
	delivery RatingUpdaterDelivery
	ratingUpdater.UnimplementedRatingUpdaterServer
}

func Register(delivery RatingUpdaterDelivery, server *grpc.Server) {
	ratingUpdater.RegisterRatingUpdaterServer(server, &RatingUpdaterServer{delivery: delivery})
}
