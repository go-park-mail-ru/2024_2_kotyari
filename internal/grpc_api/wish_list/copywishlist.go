package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
)

// CopyWishlist todo добавить имя
func (wlg *WishlistGrpc) CopyWishlist(ctx context.Context, in *wishlistgrpc.CopyWishlistRequest) (*wishlistgrpc.CopyWishlistResponse, error) {

	link, err := wlg.manager.CopyWishList(ctx, in.GetSourceLink(), in.GetTargetUserId())
	if err != nil {
		return nil, err
	}

	return &wishlistgrpc.CopyWishlistResponse{
		NewLink: link,
	}, nil
}
