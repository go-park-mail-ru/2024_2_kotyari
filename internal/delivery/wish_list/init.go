package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

type productsGetter interface {
	GetProductsByIDs(ctx context.Context, ids []uint32) ([]model.ProductCatalog, error)
}

type WishlistDelivery struct {
	client wishlistgrpc.WishlistServiceClient

	errResolver    errs.GetErrorCode
	log            *slog.Logger
	productsGetter productsGetter
}

func NewWishlistDelivery(client wishlistgrpc.WishlistServiceClient, productsGetter productsGetter, log *slog.Logger, errResolver errs.GetErrorCode) *WishlistDelivery {
	return &WishlistDelivery{
		client:         client,
		errResolver:    errResolver,
		log:            log,
		productsGetter: productsGetter,
	}
}
