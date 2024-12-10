package promocodes

import (
	"log/slog"
	"net/http"

	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *PromoCodesGRPC) GetUserPromoCodes(w http.ResponseWriter, r *http.Request) {
	userID, ok := utils.GetContextSessionUserID(r.Context())
	if !ok {
		utils.WriteErrorJSON(w, http.StatusUnauthorized, errs.UserNotAuthorized)

		return
	}

	newCtx, err := utils.AddMetadataRequestID(r.Context())
	if err != nil {
		p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] Failed to imbue ctx with request id",
			slog.String("error", err.Error()))

		utils.WriteErrorJSONByError(w, err, p.errorHandler)

		return
	}

	promoCodes, err := p.client.GetUserPromoCodes(newCtx, &promocodes.GetUserPromoCodesRequest{UserId: userID})

	grpcErr, ok := status.FromError(err)
	if err != nil {
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] No promo codes for this user",
					slog.String("error", err.Error()))

				utils.WriteErrorJSONByError(w, errs.NoPromoCodesForUser, p.errorHandler)

				return
			case codes.Unavailable:
				p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] Service unavailable",
					slog.String("error", err.Error()))

				utils.WriteErrorJSONByError(w, errs.FailedToRetrievePromoCodes, p.errorHandler)

				return

			default:
				p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] unexpected error",
					slog.String("error", err.Error()))

				utils.WriteErrorJSONByError(w, errs.InternalServerError, p.errorHandler)

				return
			}
		}

		p.log.Error("[PromoCodesGetterGRPC.GetPromoCode] Failed to retrieve Error",
			slog.String("error", err.Error()))

		utils.WriteErrorJSONByError(w, errs.InternalServerError, p.errorHandler)

		return
	}

	var promoCodesResp []PromoCodesResponseDTO

	for _, promoCode := range promoCodes.Promocodes {
		promoCodesResp = append(promoCodesResp, promoCodeFromGrpc(promoCode))
	}

	utils.WriteJSON(w, http.StatusOK, promoCodesFromDTOSlice(promoCodesResp))
}
