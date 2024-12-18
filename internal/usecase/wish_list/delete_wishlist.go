package wish_list

import "context"

func (wlu *WishListUsecase) DeleteWishlist(ctx context.Context, userID uint32, link string) error {
	err := wlu.wishListRepo.DeleteWishlist(ctx, userID, link)
	if err != nil {
		return err
	}

	err = wlu.wishListLinkRepo.DeleteWishListLink(ctx, link)
	if err != nil {
		return err
	}

	return nil
}
