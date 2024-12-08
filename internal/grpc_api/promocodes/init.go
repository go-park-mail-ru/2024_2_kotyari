package promocodes

import (
	"context"
	"log/slog"

	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type promoCodesGetter interface {
	GetUserPromoCodes(ctx context.Context, userID uint32) ([]model.PromoCode, error)
	GetPromoCode(ctx context.Context, userID uint32, promoCodeName string) (model.PromoCode, error)
}

type PromoCodesGRPC struct {
	promocodes.UnimplementedPromoCodesServer
	promoCodesGetter promoCodesGetter
	log              *slog.Logger
}

func NewPromoCodesGRPC(promoCodesGetter promoCodesGetter, logger *slog.Logger) *PromoCodesGRPC {
	return &PromoCodesGRPC{
		promoCodesGetter: promoCodesGetter,
		log:              logger,
	}
}
