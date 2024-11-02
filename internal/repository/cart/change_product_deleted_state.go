package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

func (cs *CartsStore) ChangeCartProductDeletedState(ctx context.Context, productID uint32, userID uint32) error {
	const query = `
		update carts 
		set is_deleted = false
		where product_id = $1 and user_id = $2;
	`

	commandTag, err := cs.db.Exec(ctx, query, productID, userID)
	if err != nil {
		cs.log.Error("[CartsStore.ChangeCartProductDeletedState] Error occurred when changing cart product", "error", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		cs.log.Error("[CartsStore.ChangeCartProductDeletedState] No rows were affected when changing cart product")

		return errs.ProductNotInCart
	}

	err = cs.ChangeCartProductCount(ctx, productID, 1, userID)
	if err != nil {
		cs.log.Error("[CartsStore.ChangeCartProductDeletedState] Error changing cart product count", slog.String("error", err.Error()))

		return err
	}

	return nil
}
