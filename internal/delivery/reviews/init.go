package reviews

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type reviewsManager interface {
	GetProductReviews(ctx context.Context, productID uint32) (model.Reviews, error)
	AddReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	UpdateReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	DeleteReview(ctx context.Context, productID uint32, userID uint32) error
}

type ReviewsHandler struct {
	reviewsManager reviewsManager
	inputValidator *utils.InputValidator
	errResolver    errs.GetErrorCode
	log            *slog.Logger
}

func NewReviewsHandler(manager reviewsManager, validator *utils.InputValidator, code errs.GetErrorCode, logger *slog.Logger) *ReviewsHandler {
	return &ReviewsHandler{
		reviewsManager: manager,
		inputValidator: validator,
		errResolver:    code,
		log:            logger,
	}
}
