package promocodes

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (p *PromoCodesStore) AddPromoCode(ctx context.Context, userID uint32, promoID uint32) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	p.log.Info("[PromoCodesStore.AddPromoCode] Started executing", slog.Any("request-id", requestID))

	const query = `
		insert into user_promocodes(user_id, promo_id)
		values ($1, $2);
	`

	commandTag, err := p.db.Exec(ctx, query, userID, promoID)
	if err != nil {
		p.log.Error("[PromoCodesStore.AddPromoCode] Error occurred inserting promo code",
			slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		p.log.Error("[PromoCodesStore.AddPromoCode] No rows were affected")

		return err
	}

	return nil
}
