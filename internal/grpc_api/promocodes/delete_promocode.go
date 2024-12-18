package promocodes

import (
	"context"
	"google.golang.org/grpc/codes"

	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (r *PromoCodesGRPC) DeletePromoCode(ctx context.Context, request *promocodes.DeletePromoCodesRequest) (*emptypb.Empty, error) {
	err := r.promoCodesRepo.DeletePromoCode(ctx, request.GetUserId(), request.GetPromoId())
	if err != nil {
		r.log.Error("[PromoCodesGRPC.DeletePromoCode] Failed to delete promocode")

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
