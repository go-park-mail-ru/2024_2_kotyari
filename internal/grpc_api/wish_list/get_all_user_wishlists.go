package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/golang/protobuf/ptypes/empty"
)

func (wlg *WishlistGrpc) GetAllUserWishlists(ctx context.Context, in *wishlistgrpc.GetAllWishlistsRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
