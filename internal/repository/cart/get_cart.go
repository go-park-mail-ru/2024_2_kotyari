package cart

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (cs *CartsStore) GetCart(ctx context.Context, userID uint32) (model.Cart, error) {
	const query = `
		select p.id, name, price, image_url, original_price, discount, p.count from products p
		join carts c on p.id = c.product_id where user_id=$1;
	`

	var cart model.Cart

	rows, err := cs.db.Query(ctx, query, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Cart{}, errs.CartDoesNotExist
		}

		return model.Cart{}, err
	}

	var product model.CartProduct
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.ImageURL, &product.)
	}
}
