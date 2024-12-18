package wish_list

import (
	"context"
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"github.com/golang/protobuf/ptypes/empty"
)

// todo добавить возвращение ссылки на лист

func (wlg *WishlistGrpc) CreateWishlist(ctx context.Context, in *wishlistgrpc.CreateWishlistRequest) (*empty.Empty, error) {
	err := wlg.manager.CreateWishList(ctx, in.GetUserId(), in.GetName())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
