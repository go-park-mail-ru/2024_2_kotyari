package wish_list

import (
	"context"
	"log/slog"

	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type wishlistManager interface {
	AddProductToWishlists(ctx context.Context, userSession uint32, links []string, productID uint32) error
	CopyWishList(ctx context.Context, sourceLink string, targetUserId uint32) (string, error)
	CreateWishList(ctx context.Context, userID uint32, name string) error
	DeleteWishlist(ctx context.Context, userID uint32, link string) error
	GetALlUserWishlists(ctx context.Context, userID uint32) ([]model.Wishlist, error)
	GetWishListByLink(ctx context.Context, link string) (model.Wishlist, uint32, error)
	RemoveFromWishlists(ctx context.Context, userID uint32, links []string, productId uint32) error
	RenameWishList(ctx context.Context, userID uint32, newName string, link string) error
}

type WishlistGrpc struct {
	wishlistgrpc.UnimplementedWishlistServiceServer
	manager wishlistManager
	log     *slog.Logger
}

func NewWishlistsGrpc(
	manager wishlistManager,
	log *slog.Logger,
) *WishlistGrpc {
	return &WishlistGrpc{
		manager: manager,
		log:     log,
	}
}
