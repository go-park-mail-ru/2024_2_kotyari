package cart

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cs *CartsStore) ChangeCartProductCount(ctx context.Context, productID uint32, count int32) error {
	userID := utils.GetContextSessionUserID(ctx)

	tx, err := cs.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		cs.log.Error("[CartStore.RemoveCartProduct] failed to start transaction", "error", slog.String("error", err.Error()))

		return errs.InternalServerError

	}

	defer func() {
		if err != nil {
			cs.log.Error("[CartStore.ChangeProductCount] an error occurred", slog.String("error", err.Error()))

			err = tx.Rollback(ctx)
			cs.log.Error("[CartStore.ChangeProductCount] Transaction error occurred", slog.String("error", err.Error()))
		}

		err = tx.Commit(ctx)
		if err != nil {
			cs.log.Error("[CartStore.ChangeProductCount] Transaction error occurred", slog.String("error", err.Error()))
		}
	}()

	const updateProductCount = `
		update products 
		set count = count - $2
		where id = $1;	
	`

	const updateCartProductCount = `
        update carts
        set count = count + $2
        where product_id = $1 and user_id = $3;
    `

	commandTag, err := cs.db.Exec(ctx, updateProductCount, productID, count)
	if err != nil {
		cs.log.Error("[CartStore.ChangeProductCount] Error changing product count", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		cs.log.Error("[CartStore.ChangeProductCount] Changing product count didn't affect any rows")

		return errs.ProductNotFound
	}

	commandTag, err = cs.db.Exec(ctx, updateCartProductCount, productID, count, userID)
	if err != nil {
		cs.log.Error("[CartStore.ChangeProductCount] Error changing product count in cart", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		cs.log.Error("[CartStore.ChangeProductCount] Changing product count in cart didn't affect any rows")

		return errs.ProductToModifyNotFoundInCart
	}

	return nil
}
