package product

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (ps *ProductsStore) GetProductCardByID(ctx context.Context, productID uint64) (model.ProductCard, error) {
	ps.log.Debug("[ ProductsStore.GetProductCardByID ] is running")

	card, err := ps.getProductInfo(ctx, productID)
	if err != nil {
		return model.ProductCard{}, err
	}

	categories, err := ps.getProductCategories(ctx, productID)
	if err != nil {
		return model.ProductCard{}, err
	}

	card.Categories = categories

	options, err := ps.getProductOptions(ctx, productID)
	if err != nil {
		return model.ProductCard{}, err
	}

	card.Options = options

	images, err := ps.getProductImages(ctx, productID)
	if err != nil {
		return model.ProductCard{}, err
	}

	card.Images = images
	card.ReviewCount = 0

	ps.log.Debug("[ ProductsStore.GetProductCardByID ] ProductBase successfully retrieved")

	return card, nil
}

func (ps *ProductsStore) getProductInfo(ctx context.Context, productID uint64) (model.ProductCard, error) {
	const query = `
    SELECT 
        p.id, p.title, p.count, 
        p.price, p.original_price, p.discount,
        p.rating,  p.description, p.characteristics::jsonb, 
        s.id, s.name, s.logo
    FROM products p
        JOIN sellers s ON p.seller_id = s.id
    WHERE p.id = $1 AND p.active = true;`

	row := ps.db.QueryRow(ctx, query, productID)

	var (
		card                model.ProductCard
		characteristicsJSON []byte
		seller              model.Seller
	)

	err := row.Scan(
		&card.ID, &card.Title, &card.Count,
		&card.Price, &card.OriginalPrice, &card.Discount,
		&card.Rating, &card.Description,
		&characteristicsJSON,
		&seller.ID, &seller.Name, &seller.Logo,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ps.log.Error("[ ProductsStore.GetProductCardByID ] ProductBase not found", slog.Uint64("productID", productID))
			return model.ProductCard{}, errs.ProductsDoesNotExists
		}

		ps.log.Error("[ ProductsStore.GetProductCardByID ] Error scanning row", slog.String("error", err.Error()))

		return model.ProductCard{}, err
	}

	err = json.Unmarshal(characteristicsJSON, &card.Characteristics)
	if err != nil {
		ps.log.Error("[ ProductsStore.GetProductCardByID ] Error decoding characteristics", slog.String("error", err.Error()))

		return model.ProductCard{}, err
	}

	card.Seller = seller

	return card, nil
}
