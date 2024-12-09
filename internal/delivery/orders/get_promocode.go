package orders

import (
	"context"
	"log/slog"

	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *PromoCodesGRPC) GetPromoCode(ctx context.Context, userID uint32, promoCodeName string) (model.PromoCode, error) {
	newCtx, err := utils.AddMetadataRequestID(ctx)
	if err != nil {
		p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] Failed to imbue ctx with request id",
			slog.String("error", err.Error()))

		return model.PromoCode{}, err
	}

	promoCodeResp, err := p.client.GetPromoCode(newCtx, &promocodes.GetPromoCodeRequest{
		UserId: userID,
		Name:   promoCodeName,
	})

	grpcErr, ok := status.FromError(err)
	if err != nil {
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] No promo code for this user",
					slog.String("error", err.Error()))

				return model.PromoCode{}, errs.NoPromoCode
			case codes.Unavailable:
				p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] Service unavailable",
					slog.String("error", err.Error()))

				return model.PromoCode{}, errs.FailedToRetrievePromoCode

			default:
				p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] unexpected error",
					slog.String("error", err.Error()))

				return model.PromoCode{}, errs.InternalServerError
			}
		}

		p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] Failed to retrieve Error",
			slog.String("error", err.Error()))

		return model.PromoCode{}, err
	}

	promoCode := promoCodeResp.GetPromoCode()

	return model.PromoCode{
		ID:     promoCode.GetId(),
		UserID: promoCode.GetUserId(),
		Name:   promoCode.GetName(),
		Bonus:  promoCode.GetBonus(),
	}, nil
}
