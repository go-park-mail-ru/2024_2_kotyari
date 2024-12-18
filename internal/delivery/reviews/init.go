package reviews

import (
	"context"
	"fmt"
	"log/slog"

	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type reviewsManager interface {
	AddReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	UpdateReview(ctx context.Context, productID uint32, userID uint32, review model.Review) error
	DeleteReview(ctx context.Context, productID uint32, userID uint32) error
	GetProductReviews(ctx context.Context, productID uint32, sortField string, sortOrder string) (model.Reviews, error)
}

type ReviewsHandler struct {
	reviewsManager  reviewsManager
	stringSanitizer utils.StringSanitizer
	errResolver     errs.GetErrorCode
	log             *slog.Logger
}

func NewReviewsHandler(manager reviewsManager, stringSanitizer utils.StringSanitizer, code errs.GetErrorCode, logger *slog.Logger) *ReviewsHandler {
	return &ReviewsHandler{
		reviewsManager:  manager,
		stringSanitizer: stringSanitizer,
		errResolver:     code,
		log:             logger,
	}
}

type RatingUpdaterGRPC struct {
	client ratingUpdater.RatingUpdaterClient
	log    *slog.Logger
}

func NewRatingUpdaterGRPC(config map[string]any, logger *slog.Logger) (*RatingUpdaterGRPC, error) {
	cfg, err := configs.ParseServiceViperConfig(config)
	if err != nil {
		slog.Error("[NewReviewsService] Failed to parse cfg",
			slog.String("error", err.Error()))

		return nil, err
	}

	ratingUpdaterConnection, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Domain, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("[NewReviewsService] Failed to establish gRPC connection",
			slog.String("error", err.Error()))

		return nil, err
	}

	client := ratingUpdater.NewRatingUpdaterClient(ratingUpdaterConnection)

	return &RatingUpdaterGRPC{
		client: client,
		log:    logger,
	}, nil
}
