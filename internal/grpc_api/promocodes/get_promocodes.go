package promocodes

import (
	"context"
	"errors"
	"log/slog"

	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *PromoCodesGRPC) GetUserPromoCodes(ctx context.Context, request *promocodes.GetUserPromoCodesRequest) (*promocodes.GetUserPromoCodesResponse, error) {
	promoCodes, err := r.promoCodesGetter.GetUserPromoCodes(ctx, request.GetUserId())
	if err != nil {
		if errors.Is(err, errs.NoPromoCodesForUser) {
			r.log.Error("[PromoCodesGRPC.GetPromoCodes] No promo codes for user",
				slog.String("error", err.Error()))

			return nil, status.Error(codes.NotFound, err.Error())
		}

		r.log.Error("[PromoCodesGRPC.GetPromoCodes] Unexpected error happened",
			slog.String("error", err.Error()))

		return nil, status.Error(codes.Internal, err.Error())
	}

	var resp promocodes.GetUserPromoCodesResponse
	resp.Promocodes = make([]*promocodes.PromoCode, 0, len(promoCodes))

	for _, promoCode := range promoCodes {
		resp.Promocodes = append(resp.Promocodes, promoCodeToGRPC(promoCode))
	}

	return &resp, nil
}
