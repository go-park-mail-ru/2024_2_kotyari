package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/golang/protobuf/ptypes/empty"
)

func (wlg *WishlistGrpc) AddProductToWishlists(ctx context.Context, in *wishlistgrpc.AddProductRequest) (*empty.Empty, error) {
	err := wlg.manager.AddProductToWishlists(ctx, in.GetUserId(), in.GetLinks(), in.GetProductId())
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
