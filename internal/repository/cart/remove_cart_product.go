package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (cs *CartsStore) RemoveCartProduct(ctx context.Context, productID uint32, count int32, userID uint32) error {
	tx, err := cs.db.BeginTx(ctx, pgx.TxOptions{
		AccessMode: pgx.ReadWrite,
	})

	if err != nil {
		cs.log.Error("[CartStore.RemoveCartProduct] failed to start transaction", "error", slog.String("error", err.Error()))

		return errs.InternalServerError

	}

	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	cs.log.Info("[CartsStore.RemoveCartProduct] Started executing", slog.Any("request-id", requestID))

	defer func() {
		if err != nil {
			cs.log.Error("[CartStore.RemoveCartProduct] an error occurred", slog.String("error", err.Error()))

			err = tx.Rollback(ctx)
			cs.log.Error("[CartStore.RemoveCartProduct] Transaction error occurred", slog.String("error", err.Error()))
		}

		err = tx.Commit(ctx)
		if err != nil {
			cs.log.Error("[CartStore.RemoveCartProduct] Transaction error occurred", slog.String("error", err.Error()))
		}
	}()

	err = cs.deleteCartProduct(ctx, productID, userID)
	if err != nil {
		cs.log.Error("[CartStore.RemoveCartProduct] Error deleting product from cart", slog.String("error", err.Error()))

		return err
	}

	err = cs.updateProductCount(ctx, productID, count)
	if err != nil {
		cs.log.Error("[CartStore.RemoveCartProduct] Error changing product count", slog.String("error", err.Error()))

		return err
	}

	return nil
}

func (cs *CartsStore) deleteCartProduct(ctx context.Context, productID uint32, userID uint32) error {
	const deleteQuery = `
		update carts
		set count = 0, is_deleted = true, is_selected = false
		where product_id = $1 and user_id = $2;
	`

	commandTag, err := cs.db.Exec(ctx, deleteQuery, productID, userID)
	if err != nil {
		cs.log.Error("[CartsStore.RemoveCartProduct] Error occurred when removing cart", "error", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		cs.log.Error("[CartsStore.RemoveCartProduct] No rows were affected when removing cart")

		return errs.ProductNotInCart
	}

	return nil
}
