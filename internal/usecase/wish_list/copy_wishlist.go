package wish_list

import (
	"context"
	"github.com/google/uuid"
)

func (wlu *WishListUsecase) CopyWishList(ctx context.Context, sourceLink string, targetUserId uint32) (string, error) {
	userId, err := wlu.wishListLinkRepo.GetUserIDFromLink(ctx, sourceLink)
	if err != nil {
		return "", err
	}

	newLink := uuid.NewString()

	err = wlu.wishListRepo.CopyWishlist(ctx, userId, sourceLink, targetUserId, newLink)
	if err != nil {
		return "", err
	}

	return newLink, nil
}
