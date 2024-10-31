package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/jackc/pgx/v5"
)

func (cs *CartsStore) RemoveCartProduct(ctx context.Context, productID uint32, count int32) error {
	const updateProductCount = `
		update products 
		set count = count - $2
		where id = $1;	
	`

	const deleteQuery = `
		delete from carts
		where product_id= $1;
	`
	tx, err := cs.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		cs.log.Error("[CartStore.RemoveCartProduct] failed to start transaction", "error", slog.String("error", err.Error()))

		return errs.InternalServerError

	}

	defer func() {
		if err != nil {
			cs.log.Error("[CartStore.RemoveCartProduct] an error occurred", slog.String("error", err.Error()))

			err = tx.Rollback(ctx)
			cs.log.Error("[CartStore.ChangeProductCount] Transaction error occurred", slog.String("error", err.Error()))
		}

		err = tx.Commit(ctx)
		if err != nil {
			cs.log.Error("[CartStore.ChangeProductCount] Transaction error occurred", slog.String("error", err.Error()))
		}
	}()

	commandTag, err := cs.db.Exec(ctx, deleteQuery, productID)
	if err != nil {
		cs.log.Error("[CartsStore.RemoveCartProduct] Error occurred when removing cart", "error", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		cs.log.Error("[CartsStore.RemoveCartProduct] No rows were affected when removing cart", "error", slog.String("error", err.Error()))

		return err
	}

	commandTag, err = cs.db.Exec(ctx, updateProductCount, productID, count)
	if err != nil {
		cs.log.Error("[CartStore.ChangeProductCount] Error changing product count", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		cs.log.Error("[CartStore.ChangeProductCount] Changing product count didn't affect any rows")

		return errs.ProductNotFound
	}

	return nil
}
