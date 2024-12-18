package wish_list

import (
	wishlistgrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/wish_list/gen"
	"google.golang.org/grpc"
)

func (wlg *WishlistGrpc) Register(grpcServer *grpc.Server) {
	wishlistgrpc.RegisterWishlistServiceServer(grpcServer, wlg)
}
