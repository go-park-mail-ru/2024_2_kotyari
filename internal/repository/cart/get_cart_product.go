package cart

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (cs *CartsStore) GetCartProduct(ctx context.Context, productID uint32, userID uint32) (model.CartProduct, error) {
	const query = `
		select c.count, c.is_selected, c.is_deleted from carts c
		join products p on p.id = c.product_id
		where p.id = $1 and c.user_id = $2;
	`

	var cartProduct model.CartProduct

	err := cs.db.QueryRow(ctx, query, productID, userID).Scan(&cartProduct.Count, &cartProduct.IsSelected, &cartProduct.IsDeleted)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			cs.log.Error("[CartStore.GetCartProduct] This cartProduct is not in cart ", slog.String("error", err.Error()))

			return model.CartProduct{}, errs.ProductNotInCart
		}

		cs.log.Error("[CartStore.GetCartProduct] Error performing query: ", slog.String("error", err.Error()))

		return model.CartProduct{}, err
	}

	return cartProduct, nil
}
