package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (cs *CartsStore) AddProduct(ctx context.Context, productID uint32, userID uint32) error {
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

	cs.log.Info("[CartsStore.AddProduct] Started executing", slog.Any("request-id", requestID))

	const query = `
		insert into carts(user_id, product_id, count, is_selected, is_deleted)
		values ($1, $2, 1, true, false)
	`

	commandTag, err := cs.db.Exec(ctx, query, userID, productID)
	if err != nil {
		cs.log.Error("[CartStore.AddProduct] Error adding product", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() != 1 {
		cs.log.Error("[CartStore.AddProduct] Adding product count didn't affect any rows")

		return err
	}

	err = cs.updateProductCount(ctx, productID, 1)
	if err != nil {
		cs.log.Error("[CartStore.ChangeProductCount] Error occurred when changing cart product count", slog.String("error", err.Error()))

		return err
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

	return nil
}
