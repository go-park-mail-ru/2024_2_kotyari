package wish_list

import (
	"context"
	"github.com/google/uuid"
)

func (wlu *WishListUsecase) CreateWishList(ctx context.Context, userID uint32, name string) error {
	link := uuid.NewString()

	err := wlu.wishListRepo.CreateWishlist(ctx, userID, name, link)
	if err != nil {
		return err
	}

	err = wlu.wishListLinkRepo.CreateLink(ctx, userID, link)
	if err != nil {
		return err
	}

	return nil
}
