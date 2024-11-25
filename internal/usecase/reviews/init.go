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
	GetProductReviewsNoLogin(ctx context.Context, productID uint32, sortField string, sortOrder string) (model.Reviews, error)
	GetProductReviewsWithLogin(ctx context.Context, productID uint32, userID uint32, sortField string, sortOrder string) (model.Reviews, error)
}

type ratingUpdater interface {
	UpdateRating(ctx context.Context, productID uint32) error
}

type ReviewsService struct {
	reviewsRepo     reviewsRepo
	stringSanitizer utils.StringSanitizer
	log             *slog.Logger
	ratingUpdater   ratingUpdater
}

func NewReviewsService(repo reviewsRepo, stringSanitizer utils.StringSanitizer, logger *slog.Logger, updater ratingUpdater) (*ReviewsService, error) {
	return &ReviewsService{
		reviewsRepo:     repo,
		stringSanitizer: stringSanitizer,
		log:             logger,
		ratingUpdater:   updater,
	}, nil
}
