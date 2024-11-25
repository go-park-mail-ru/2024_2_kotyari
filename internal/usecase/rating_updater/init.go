package rating_updater

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type RatingUpdaterRepository interface {
	UpdateProductRating(ctx context.Context, productID uint32, newRating float32) error
}

type ReviewsGetter interface {
	GetProductReviewsNoLogin(ctx context.Context, productID uint32, sortField string, sortOrder string) (model.Reviews, error)
}

type RatingUpdaterService struct {
	repository    RatingUpdaterRepository
	reviewsGetter ReviewsGetter
	log           *slog.Logger
}

func NewRatingUpdateService(repository RatingUpdaterRepository, reviewsFetcher ReviewsGetter, logger *slog.Logger) *RatingUpdaterService {
	return &RatingUpdaterService{
		repository:    repository,
		reviewsGetter: reviewsFetcher,
		log:           logger,
	}
}
