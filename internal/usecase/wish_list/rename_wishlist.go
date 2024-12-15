package wish_list

import "context"

func (wlu *WishListUsecase) RenameWishList(ctx context.Context, userID uint32, newName string, link string) error {
	return wlu.wishListRepo.RenameWishlist(ctx, userID, newName, link)
}
