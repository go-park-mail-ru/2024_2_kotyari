package delivery

import (
	"context"
)

type RatingUpdaterRepository interface {
	UpdateProductRating(ctx context.Context, productID uint32, newRating float32) error
}

type RatingUpdaterGRPCHandler struct {
	repository RatingUpdaterRepository
}
