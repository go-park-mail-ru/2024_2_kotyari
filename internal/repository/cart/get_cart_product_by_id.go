package cart

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cs *CartsStore) GetCartProductByID(ctx context.Context, productID uint32) (model.BaseProduct, error) {
	userID := utils.GetContextSessionUserID(ctx)

	// Мб поменять на select count from products where id=$1;
	const query = `
		select p.id, p.count from products p 
		join carts c on p.id = c.product_id where user_id=$1 and product_id=$2;
	`

	var product model.BaseProduct

	err := cs.db.QueryRow(ctx, query, userID, productID).Scan(&product.ID, &product.Count)
	if err != nil {
		cs.log.Error("[CartStore.GetProductQuantity] Error performing query: ", slog.String("error", err.Error()))

		return model.BaseProduct{}, err
	}

	return product, nil
}
