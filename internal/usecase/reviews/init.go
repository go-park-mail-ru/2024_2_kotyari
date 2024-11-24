package reviews

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type reviewsRepo interface {
	GetReview(ctx context.Context, productID uint32, userID uint32) (model.Review, error)
	AddReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	UpdateReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	DeleteReview(ctx context.Context, productID uint32, userID uint32) error
}

type ratingUpdater interface {
	UpdateRating(ctx context.Context, productID uint32) error
}

type ReviewsService struct {
	reviewsRepo    reviewsRepo
	inputValidator *utils.InputValidator
	log            *slog.Logger
	ratingUpdater  ratingUpdater
}

func NewReviewsService(repo reviewsRepo, validator *utils.InputValidator, logger *slog.Logger, updater ratingUpdater) (*ReviewsService, error) {
	return &ReviewsService{
		reviewsRepo:    repo,
		inputValidator: validator,
		log:            logger,
		ratingUpdater:  updater,
	}, nil
}
