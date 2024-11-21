package app

import (
	"context"
	"log/slog"

	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"google.golang.org/grpc"
)

type RatingUpdaterManager interface {
	UpdateProductRating(ctx context.Context, productID uint32) error
}

type RatingUpdaterServer struct {
	manager RatingUpdaterManager
	log     *slog.Logger
	ratingUpdater.UnimplementedRatingUpdaterServer
}

func Register(manager RatingUpdaterManager, logger *slog.Logger, server *grpc.Server) {
	ratingUpdater.RegisterRatingUpdaterServer(server, &RatingUpdaterServer{manager: manager, log: logger})
}
