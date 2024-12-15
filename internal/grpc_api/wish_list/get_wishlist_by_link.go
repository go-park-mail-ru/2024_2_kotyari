package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"google.golang.org/grpc"
)

func (wlg *WishlistGrpc) GetWishlistByLink(ctx context.Context, in *wishlistgrpc.GetWishlistByLinkRequest, opts ...grpc.CallOption) (*wishlistgrpc.GetWishlistByLinkResponse, error) {

	return nil, nil
}
