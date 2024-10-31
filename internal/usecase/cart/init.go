package cartServiceLib

import (
	"context"
	"log/slog"
)

type cartRepository interface {
	GetCartProductCount(ctx context.Context, productID uint32) (uint32, error)
	ChangeCartProductCount(ctx context.Context, productID uint32, count int32) error
}

type CartManager struct {
	cartRepository cartRepository
	log            *slog.Logger
}

func NewCartManager(repository cartRepository, logger *slog.Logger) *CartManager {
	return &CartManager{
		cartRepository: repository,
		log:            logger,
	}
}
