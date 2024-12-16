package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/golang/protobuf/ptypes/empty"
)

func (wlg *WishlistGrpc) RenameWishlist(ctx context.Context, in *wishlistgrpc.RenameWishlistRequest) (*empty.Empty, error) {
	err := wlg.manager.RenameWishList(ctx, in.GetUserId(), in.GetNewName(), in.GetLink())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
