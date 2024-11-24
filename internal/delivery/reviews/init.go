package reviews

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

type reviewsManager interface {
	AddReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	UpdateReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	DeleteReview(ctx context.Context, productID uint32, userID uint32) error
}

type reviewsGetter interface {
	GetProductReviewsNoLogin(ctx context.Context, productID uint32, sortField string, sortOrder string) (model.Reviews, error)
	GetProductReviewsWithLogin(ctx context.Context, productID uint32, userID uint32, sortField string, sortOrder string) (model.Reviews, error)
}

type ReviewsHandler struct {
	reviewsManager reviewsManager
	reviewsGetter  reviewsGetter
	inputValidator *utils.InputValidator
	errResolver    errs.GetErrorCode
	log            *slog.Logger
}

func NewReviewsHandler(manager reviewsManager, reviewsGetter reviewsGetter, validator *utils.InputValidator, code errs.GetErrorCode, logger *slog.Logger) *ReviewsHandler {
	return &ReviewsHandler{
		reviewsManager: manager,
		reviewsGetter:  reviewsGetter,
		inputValidator: validator,
		errResolver:    code,
		log:            logger,
	}
}
