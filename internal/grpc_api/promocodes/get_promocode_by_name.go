package promocodes

import (
	"context"
	"errors"
	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (r *PromoCodesGRPC) GetPromoCode(ctx context.Context, request *promocodes.GetPromoCodeRequest) (*promocodes.GetPromoCodeResponse, error) {
	promoCode, err := r.promoCodesGetter.GetPromoCode(ctx, request.GetUserId(), request.GetName())
	if err != nil {
		if errors.Is(err, errs.NoPromoCode) {
			r.log.Error("[PromoCodesGRPC.GetPromoCode] No promo code", slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, err.Error())
		}

		r.log.Error("[PromoCodesGRPC.GetPromoCode] Unexpected error happened", slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &promocodes.GetPromoCodeResponse{
		PromoCode: promoCodeToGRPC(promoCode),
	}, nil
}
