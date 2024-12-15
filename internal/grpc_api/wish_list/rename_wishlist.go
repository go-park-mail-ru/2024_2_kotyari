package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/golang/protobuf/ptypes/empty"
)

func (wlg *WishlistGrpc) RenameWishlist(ctx context.Context, in *wishlistgrpc.RenameWishlistRequest) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}
