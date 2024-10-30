package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cs *CartsStore) GetCartProductQuantity(ctx context.Context, productID uint32) (uint32, error) {
	userID := utils.GetContextSessionUserID(ctx)

	// Мб поменять на select count from products where id=$1;
	const query = `
		select p.count from products p 
		join carts c on p.id = c.product_id where user_id=$1 and product_id=$2;
	`

	var count uint32

	err := cs.db.QueryRow(ctx, query, userID, productID).Scan(&count)
	if err != nil {
		cs.log.Error("[CartStore.GetProductQuantity] Error performing query: ", slog.String("error", err.Error()))

		return 0, err
	}

	return count, nil
}
