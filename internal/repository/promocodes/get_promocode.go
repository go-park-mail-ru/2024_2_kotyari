package promocodes

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (p *PromoCodesStore) GetPromoCode(ctx context.Context, userID uint32, promoCodeName string) (model.PromoCode, error) {
	//requestID, err := utils.GetContextRequestID(ctx)
	//if err != nil {
	//	p.log.Error("[PromoCodesStore.GetPromoCode] Failed to get request id", slog.String("error", err.Error()))
	//
	//	return model.PromoCode{}, err
	//}
	//
	//p.log.Info("[PromoCodesStore.GetPromoCode] Started executing", slog.Any("request-id", requestID))
	//

	const query = `
		select p.id, up.user_id, p.name, p.bonus, p.updated_at, p.created_at
		from promocodes p
		join user_promocodes up on p.id = up.promo_id
		where up.user_id = $1 and p.name = $2;
	`

	var promoCode PromoCodesDTO

	err := p.db.QueryRow(ctx, query, userID, promoCodeName).Scan(
		&promoCode.ID,
		&promoCode.UserID,
		&promoCode.Name,
		&promoCode.Bonus,
		&promoCode.UpdatedAt,
		&promoCode.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.log.Error("[PromoCodesStore.GetPromoCode] No rows", slog.String("error", err.Error()))

			return model.PromoCode{}, errs.NoPromoCode
		}

		p.log.Error("[PromoCodesStore.GetPromoCode] Unexpected error", slog.String("error", err.Error()))

		return model.PromoCode{}, err
	}

	return promoCode.ToModel(), nil
}
