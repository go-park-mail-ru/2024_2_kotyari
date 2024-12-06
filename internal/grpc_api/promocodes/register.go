package promocodes

import (
	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"google.golang.org/grpc"
)

func (r *PromoCodesGRPC) Register(server *grpc.Server) {
	promocodes.RegisterPromoCodesServer(server, r)
}
