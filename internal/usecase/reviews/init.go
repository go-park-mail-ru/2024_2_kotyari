package reviews

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type reviewsRepo interface {
	GetProductReviews(ctx context.Context, productID uint32) (model.Reviews, error)
	GetReview(ctx context.Context, productID uint32, userID uint32) (model.Review, error)
	AddReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	UpdateReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	DeleteReview(ctx context.Context, productID uint32, userID uint32) error
}

type ReviewsService struct {
	reviewsRepo    reviewsRepo
	inputValidator *utils.InputValidator
	log            *slog.Logger
}

func NewReviewsService(repo reviewsRepo, validator *utils.InputValidator, logger *slog.Logger) *ReviewsService {
	return &ReviewsService{
		reviewsRepo:    repo,
		inputValidator: validator,
		log:            logger,
	}
}
