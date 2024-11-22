package usecase

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type RatingUpdaterRepository interface {
	UpdateProductRating(ctx context.Context, productID uint32, newRating float32) error
}

type ReviewsFetcher interface {
	GetProductReviews(ctx context.Context, productID uint32, sortField string, sortOrder string) (model.Reviews, error)
}

type RatingUpdaterService struct {
	repository     RatingUpdaterRepository
	reviewsFetcher ReviewsFetcher
	log            *slog.Logger
}

func NewRatingUpdateService(repository RatingUpdaterRepository, reviewsFetcher ReviewsFetcher, logger *slog.Logger) *RatingUpdaterService {
	return &RatingUpdaterService{
		repository:     repository,
		reviewsFetcher: reviewsFetcher,
		log:            logger,
	}
}
