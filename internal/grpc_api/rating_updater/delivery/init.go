package delivery

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

type RatingUpdaterRepository interface {
	UpdateProductRating(ctx context.Context, productID uint32, newRating float32) error
}

type RatingUpdaterHandler struct {
	repository  RatingUpdaterRepository
	log         *slog.Logger
	errResolver errs.GetErrorCode
}

func NewRatingUpdaterHandler(repository RatingUpdaterRepository, logger *slog.Logger, code errs.GetErrorCode) *RatingUpdaterHandler {
	return &RatingUpdaterHandler{
		repository:  repository,
		log:         logger,
		errResolver: code,
	}
}
