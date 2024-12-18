package cart

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const basicUrl = "/catalog/product/"

func mapRowToCartProduct(
	id pgtype.Int4,
	title pgtype.Text,
	price pgtype.Uint32,
	image pgtype.Text,
	weight pgtype.Float4,
	quantity pgtype.Uint32,
) model.CartProductForOrder {
	product := model.CartProductForOrder{}

	if id.Valid {
		product.URL = basicUrl + strconv.Itoa(int(id.Int32))
	} else {
		product.URL = ""
	}
	product.Title = title.String
	product.Price = price.Uint32
	product.Image = image.String
	product.Weight = weight.Float32
	product.Quantity = quantity.Uint32
	product.DeliveryDate = time.Now().Add(72 * time.Hour)

	return product
}

func (cs *CartsStore) GetSelectedFromCart(ctx context.Context, userID uint32) (*model.CartProductsForOrderWithUser, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	cs.log.Info("[CartsStore.GetSelectedFromCart] Started executing", slog.Any("request-id", requestID))

	const query = `
		SELECT p.id, p.title, p.price, p.image_url, p.weight, c.count, c.delivery_date, u.username, u.preferred_payment_method,
           a.address
		FROM users u
			LEFT JOIN carts c ON u.id = c.user_id AND c.is_deleted = false and c.is_selected = true
			LEFT JOIN products p ON p.id = c.product_id
			LEFT JOIN addresses a ON u.id = a.user_id
		WHERE u.id=$1;
	`

	rows, err := cs.db.Query(ctx, query, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			cs.log.Error("[CartStore.GetCartProducts] No rows found")
			return nil, errors.New("cart does not exist")
		}
		cs.log.Error("[CartStore.GetCartProducts] Unexpected error occurred")
		return nil, err
	}

	var (
		username               string
		preferredPaymentMethod string
		address                pgtype.Text
	)

	var products []model.CartProductForOrder
	for rows.Next() {
		var (
			id           pgtype.Int4
			title        pgtype.Text
			price        pgtype.Uint32
			image        pgtype.Text
			weight       pgtype.Float4
			quantity     pgtype.Uint32
			deliveryDate pgtype.Timestamptz
		)

		err = rows.Scan(
			&id,
			&title,
			&price,
			&image,
			&weight,
			&quantity,
			&deliveryDate,
			&username,
			&preferredPaymentMethod,
			&address,
		)
		if err != nil {
			cs.log.Error("[CartStore.GetCartProducts] Error scanning row")
			return nil, err
		}

		product := mapRowToCartProduct(id, title, price, image, weight, quantity)
		products = append(products, product)
	}

	cart := model.CartProductsForOrderWithUser{
		UserName:               username,
		PreferredPaymentMethod: preferredPaymentMethod,
		Items:                  products,
		Address: model.Address{
			Text: address.String,
		},
	}

	return &cart, nil
}
