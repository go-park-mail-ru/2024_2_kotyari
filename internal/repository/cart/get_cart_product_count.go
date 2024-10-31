package cart

import (
	"context"
	"log/slog"
)

func (cs *CartsStore) GetCartProductCount(ctx context.Context, productID uint32) (uint32, error) {

	// Мб поменять на select count from products where id=$1;
	const query = `
		select count from products 
		where id = $1;
	`

	var count uint32

	err := cs.db.QueryRow(ctx, query, productID).Scan(&count)
	if err != nil {
		cs.log.Error("[CartStore.GetCartProductCount] Error performing query: ", slog.String("error", err.Error()))

		return 0, err
	}

	return count, nil
}
