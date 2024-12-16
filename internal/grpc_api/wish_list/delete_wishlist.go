package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/golang/protobuf/ptypes/empty"
)

func (wlg *WishlistGrpc) DeleteWishlist(ctx context.Context, in *wishlistgrpc.DeleteWishlistRequest) (*empty.Empty, error) {
	err := wlg.manager.DeleteWishlist(ctx, in.GetUserId(), in.GetLink())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
