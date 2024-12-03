package promocodes

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (p *PromoCodesStore) DeletePromoCode(ctx context.Context, userID uint32, promoCode model.PromoCode) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	p.log.Info("[PromoCodesStore.RemovePromoCode] Started executing", slog.Any("request-id", requestID))

	const query = `
		delete from promo_codes
		where user_id = $1 and name = $2;
	`

	commandTag, err := p.db.Exec(ctx, query, userID, promoCode.Name)
	if err != nil {
		p.log.Error("[PromoCodesStore.RemovePromoCode] Error executing query", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		p.log.Error("[PromoCodesStore.RemovePromoCode] No rows were affected")

		return err
	}

	return nil
}
