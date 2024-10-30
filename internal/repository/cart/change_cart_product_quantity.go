package cart

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"log/slog"
)

func (cs *CartsStore) ChangeCartProductQuantity(ctx context.Context, productID uint32, count int32) (uint32, error) {
	userID := utils.GetContextSessionUserID(ctx)

	const query = `
        update products p
        set count = p.count + $1
        from carts c
        where p.id = c.product_id and c.user_id = $2 and p.id = $3 returning p.count;
    `

	var newProductCount uint32

	err := cs.db.QueryRow(ctx, query, count, userID, productID).Scan(&newProductCount)
	if err != nil {
		cs.log.Error("[CartStore.ChangeProductQuantity] Error changing product quantity", slog.String("error", err.Error()))

		return 0, err
	}

	return newProductCount, nil
}
