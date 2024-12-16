package wish_list

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

func (wlu *WishListUsecase) AddProductToWishlists(ctx context.Context, userSession uint32, links []string, productID uint32) error {
	if links == nil || len(links) == 0 {
		return nil
	}

	userID, err := wlu.wishListLinkRepo.GetUserIDFromLink(ctx, links[0])
	if err != nil {
		return err
	}

	if userID != userSession {
		return errs.ErrNotPermitted
	}

	err = wlu.wishListRepo.AddProductToWishlists(ctx, userID, links, productID)
	if err != nil {
		return err
	}

	return nil
}
