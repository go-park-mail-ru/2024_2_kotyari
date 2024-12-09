package orders

import (
	"context"
	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (p *PromoCodesGRPC) DeletePromoCode(ctx context.Context, userID uint32, promoID uint32) error {
	newCtx, err := utils.AddMetadataRequestID(ctx)
	if err != nil {
		p.log.Error("[PromoCodesGRPC.DeletePromoCode] Failed to imbue ctx with request id",
			slog.String("error", err.Error()))

		return err
	}

	_, err = p.client.DeletePromoCode(newCtx, &promocodes.DeletePromoCodesRequest{
		UserId:  userID,
		PromoId: promoID,
	})

	grpcErr, ok := status.FromError(err)
	if err != nil {
		if ok {
			switch grpcErr.Code() {
			case codes.Unavailable:
				p.log.Error("[PromoCodesGRPC.DeletePromoCode] Service unavailable",
					slog.String("error", err.Error()))

				return err

			default:
				p.log.Error("[PromoCodesGRPC.DeletePromoCode] unexpected error",
					slog.String("error", err.Error()))

				return errs.InternalServerError
			}
		}

		p.log.Error("[PromoCodesGRPC.GetPromoCode] Failed to retrieve Error",
			slog.String("error", err.Error()))

		return err
	}

	return nil
}
