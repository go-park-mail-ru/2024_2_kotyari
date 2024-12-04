package promocodes

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (p *PromoCodesStore) GetUserPromoCodes(ctx context.Context, userID uint32) ([]model.PromoCode, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	p.log.Info("[PromoCodesStore.GetUserPromoCodes] Started executing", slog.Any("request-id", requestID))

	const query = `
		select p.id, up.user_id, p.bonus, p.updated_at, p.created_at
		from promocodes p
		join user_promocodes up on p.id = up.promo_id
		where up.user_id = $1
	`

	rows, err := p.db.Query(ctx, query, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.log.Error("[PromoCodesStore.GetUserPromoCodes] No promo codes")

			return nil, errs.NoPromoCodesForUser
		}

		p.log.Error("[PromoCodesStore.GetUserPromoCodes] Unexpected error", slog.String("error", err.Error()))

		return nil, err
	}

	promoCodes, err := pgx.CollectRows(rows, pgx.RowToStructByName[PromoCodesDTO])
	if err != nil {
		p.log.Error("[PromoCodesStore.GetUserPromoCodes] Error getting promo codes",
			slog.String("error", err.Error()))

		return nil, err
	}

	return PromoCodesToModelSlice(promoCodes), nil
}
