package wish_list

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (wlu *WishListUsecase) GetALlUserWishlists(ctx context.Context, userID uint32) ([]model.Wishlist, error) {
	return wlu.wishListRepo.GetALlUserWishlists(ctx, userID)
}
