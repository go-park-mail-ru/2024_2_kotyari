package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

func (cs *CartsStore) ChangeCartProductSelectedState(ctx context.Context, productID uint32, userID uint32, isSelected bool) error {
	const query = `
		update carts 
		set is_selected = $3
		where product_id = $1 and user_id = $2;
	`

	commandTag, err := cs.db.Exec(ctx, query, productID, userID, isSelected)
	if err != nil {
		cs.log.Error("[CartsStore.ChangeCartProductSelectedState] Error occurred when changing cart product", "error", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		cs.log.Error("[CartsStore.ChangeCartProductSelectedState] No rows were affected when changing cart product")

		return errs.ProductNotInCart
	}

	return nil
}
