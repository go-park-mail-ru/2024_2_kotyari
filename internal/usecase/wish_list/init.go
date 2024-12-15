package wish_list

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"
)

type wishListRepo interface {
	AddProductToWishlists(ctx context.Context, userID uint32, links []string, productID uint32) error
	CopyWishlist(ctx context.Context, sourceUserID uint32, sourceLink string, targetUserID uint32, newLink string) error
	CreateWishlist(ctx context.Context, userID uint32, name string, link string) error
	DeleteWishlist(ctx context.Context, userID uint32, link string) error
	GetALlUserWishlists(ctx context.Context, userID uint32) ([]model.Wishlist, error)
	GetWishListByLink(ctx context.Context, userID uint32, link string) (model.Wishlist, error)
	RemoveProductFromWishlists(ctx context.Context, userID uint32, links []string, productID uint32) error
	RenameWishlist(ctx context.Context, userID uint32, newName string, link string) error
}

type wishListLinkRepo interface {
	CreateLink(ctx context.Context, userID uint32, link string) error
	DeleteWishListLink(ctx context.Context, link string) error
	GetUserIDFromLink(ctx context.Context, link string) (uint32, error)
}

type WishListUsecase struct {
	wishListRepo     wishListRepo
	wishListLinkRepo wishListLinkRepo

	log *slog.Logger
}

func NewWushlistUsecase(wishListRepo wishListRepo, wishListLinkRepo wishListLinkRepo, log *slog.Logger) *WishListUsecase {
	return &WishListUsecase{
		wishListRepo:     wishListRepo,
		wishListLinkRepo: wishListLinkRepo,
		log:              log,
	}
}
