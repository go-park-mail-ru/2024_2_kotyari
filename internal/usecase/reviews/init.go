package reviews

import (
	"context"
	ratingUpdater "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/rating_updater/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

type ReviewsService struct {
	reviewsRepo    reviewsRepo
	inputValidator *utils.InputValidator
	log            *slog.Logger
	client         ratingUpdater.RatingUpdaterClient
}

func NewReviewsService(repo reviewsRepo, validator *utils.InputValidator, logger *slog.Logger) (*ReviewsService, error) {
	ratingUpdaterConnection, err := grpc.NewClient("rating_updater_go:8004",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logger.Error("[NewReviewsService] Failed to establish gRPC connection",
			slog.String("error", err.Error()))

		return nil, err
	}

	client := ratingUpdater.NewRatingUpdaterClient(ratingUpdaterConnection)

	return &ReviewsService{
		reviewsRepo:    repo,
		inputValidator: validator,
		log:            logger,
		client:         client,
	}, nil
}
