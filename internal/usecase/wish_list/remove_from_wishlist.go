package wish_list

import "context"

func (wlu *WishListUsecase) RemoveFromWishlists(ctx context.Context, userID uint32, links []string, productId uint32) error {
	return wlu.wishListRepo.RemoveProductFromWishlists(ctx, userID, links, productId)
}
