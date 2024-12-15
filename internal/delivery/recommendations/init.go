package recommendations

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type RecRepo interface {
	GetRecommendations(ctx context.Context, productId uint32) ([]model.ProductCatalog, error)
}

type RecDelivery struct {
	log         *slog.Logger
	errResolver errs.GetErrorCode
	repo        RecRepo
}

func NewRecHandler(
	logger *slog.Logger,
	errResolver errs.GetErrorCode,
	repo RecRepo,
) *RecDelivery {

	return &RecDelivery{
		log:         logger,
		errResolver: errResolver,
		repo:        repo,
	}
}
