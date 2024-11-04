package cart

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (cs *CartsStore) GetCart(ctx context.Context, userID uint32, deliveryDate time.Time) (model.Cart, error) {
	const query = `
		select c.id, p.id, title, price, description, image_url, original_price, discount, c.count, c.is_selected from products p
		join carts c on p.id = c.product_id where user_id=$1 and c.is_deleted = false and c.count>0;
	`

	var cart model.Cart

	rows, err := cs.db.Query(ctx, query, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			cs.log.Error("[CartStore.GetCart] Err no rows")

			return model.Cart{}, errs.CartDoesNotExist
		}

		cs.log.Error("[CartStore.GetCart] Unexpected error occurred: ", slog.String("error", err.Error()))

		return model.Cart{}, err
	}

	for rows.Next() {
		var product model.CartProduct
		err = rows.Scan(
			&cart.ID,
			&product.ID,
			&product.Title,
			&product.Price,
			&product.Description,
			&product.ImageURL,
			&product.OriginalPrice,
			&product.Discount,
			&product.Count,
			&product.IsSelected)
		if err != nil {
			cs.log.Error("[CartStore.GetCart] Error parsing rows: ", slog.String("error", err.Error()))

			return model.Cart{}, err
		}

		cart.Products = append(cart.Products, product)
	}

	if len(cart.Products) == 0 {
		return model.Cart{}, errs.EmptyCart
	}

	cart.UserID = userID
	cart.DeliveryDate = deliveryDate

	return cart, err
}
