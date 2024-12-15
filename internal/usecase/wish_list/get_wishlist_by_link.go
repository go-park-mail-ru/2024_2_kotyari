package wish_list

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (wlu *WishListUsecase) GetWishListByLink(ctx context.Context, link string) (model.Wishlist, error) {
	userID, err := wlu.wishListLinkRepo.GetUserIDFromLink(ctx, link)
	if err != nil {
		return model.Wishlist{}, err
	}

	return wlu.wishListRepo.GetWishListByLink(ctx, userID, link)
}
