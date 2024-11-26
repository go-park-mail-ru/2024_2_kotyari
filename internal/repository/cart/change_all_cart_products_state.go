package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cs *CartsStore) ChangeAllCartProductsState(ctx context.Context, userID uint32, isSelected bool) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	cs.log.Info("[CartsStore.ChangeAllCartProductsState] Started executing", slog.Any("request-id", requestID))

	const query = `
		update carts
		set is_selected = $2
		where user_id = $1 and is_deleted = false;
	`

	commandTag, err := cs.db.Exec(ctx, query, userID, isSelected)
	if err != nil {
		cs.log.Error("[CartsStore.SelectAllCartProducts] Error occurred when changing cart product", "error", slog.String("error", err.Error()))

		return err
	}

	if commandTag.RowsAffected() == 0 {
		cs.log.Error("[CartsStore.SelectAllCartProducts] No rows were affected when changing cart product")

		return errs.EmptyCart
	}

	return nil
}
