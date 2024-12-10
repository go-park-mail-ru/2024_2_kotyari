package cart

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (cs *CartsStore) GetCartProductCount(ctx context.Context, userID uint32, productID uint32) (uint32, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return 0, err
	}

	cs.log.Info("[CartsStore.GetCartProductCount] Started executing", slog.Any("request-id", requestID))

	const query = `
		select c.count
		from carts c
		where c.user_id = $1 and c.product_id = $2;
	`

	var productCount uint32

	err = cs.db.QueryRow(ctx, query, userID, productID).Scan(&productCount)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			cs.log.Error("[CartsStore.GetCartProductCount] Error no rows")

			return 0, errs.ProductNotInCart
		}

		cs.log.Error("[CartsStore.GetCartProductCount] Unexpected error",
			slog.String("error", err.Error()))

		return 0, err
	}

	return productCount, nil
}
