package rating_updater

import (
	"context"
	"log/slog"

	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
)

type RatingUpdaterManager interface {
	UpdateProductRating(ctx context.Context, productID uint32) error
}

type RatingUpdaterGRPC struct {
	manager RatingUpdaterManager
	log     *slog.Logger
	ratingUpdater.UnimplementedRatingUpdaterServer
}

func NewRatingUpdaterGRPC(manager RatingUpdaterManager, logger *slog.Logger) *RatingUpdaterGRPC {
	return &RatingUpdaterGRPC{
		manager: manager,
		log:     logger,
	}
}
