package promocodes

import (
	"context"
	"log/slog"

	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type promoCodesRepo interface {
	GetUserPromoCodes(ctx context.Context, userID uint32) ([]model.PromoCode, error)
	GetPromoCode(ctx context.Context, userID uint32, promoCodeName string) (model.PromoCode, error)
	DeletePromoCode(ctx context.Context, userID uint32, promoID uint32) error
}

type PromoCodesGRPC struct {
	promocodes.UnimplementedPromoCodesServer
	promoCodesRepo promoCodesRepo
	log            *slog.Logger
}

func NewPromoCodesGRPC(promoCodesRepo promoCodesRepo, logger *slog.Logger) *PromoCodesGRPC {
	return &PromoCodesGRPC{
		promoCodesRepo: promoCodesRepo,
		log:            logger,
	}
}
