package cart

import (
	"context"
	"log/slog"
)

func (cs *CartsStore) AddProduct(ctx context.Context, productID uint32, userID uint32) error {
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

	return nil
}
