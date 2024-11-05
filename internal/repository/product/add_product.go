package product

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
)

func (ps *ProductsStore) AddProduct(ctx context.Context, card model.ProductCard) error {
	ps.log.Debug("[ ProductsStore.AddProduct ] Adding product")

	tx, err := ps.db.Begin(ctx)
	if err != nil {
		ps.log.Error("[ ProductsStore.AddProduct ] Error starting transaction", slog.String("error", err.Error()))
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		if err != nil {
			err = tx.Rollback(context.TODO())
			if err != nil {
				ps.log.Error("[ ProductStore.AddProduct ] Error rolling back transaction ] ")
			}
		} else {
			err = tx.Commit(context.TODO())
			if err != nil {
				ps.log.Error("[ ProductStore.AddProduct ] Error committing transaction ] ")
			}
		}
	}(tx, ctx)

	var productID uint64
	insertProductQuery := `
		INSERT INTO products (
			seller_id, count, price, original_price, discount, title, description, rating, characteristics, image_url
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9::jsonb, $10)
		RETURNING id;
	`
	mainImageUrl := ""
	if len(card.Images) > 0 {
		mainImageUrl = card.Images[0].Url
	}

	characteristicsJSON, err := json.Marshal(card.Characteristics)
	if err != nil {
		ps.log.Error("[ ProductsStore.AddProduct ] Error marshalling characteristics", slog.String("error", err.Error()))
		return err
	}

	err = tx.QueryRow(ctx, insertProductQuery,
		card.Seller.ID, card.Count, card.Price,
		card.OriginalPrice, card.Discount, card.Title,
		card.Description, card.Rating, characteristicsJSON,
		mainImageUrl).
		Scan(&productID)
	if err != nil {
		ps.log.Error("[ ProductsStore.AddProduct ] Error inserting product", slog.String("error", err.Error()))
		return err
	}

	// Insert categories
	insertCategoriesQuery := `
		INSERT INTO product_categories (product_id, category_id)
		SELECT $1, unnest($2::bigint[]);`

	categoryIDs := make([]uint32, len(card.Categories))
	for i, category := range card.Categories {
		categoryIDs[i] = category.ID
	}

	_, err = tx.Exec(ctx, insertCategoriesQuery, productID, categoryIDs)
	if err != nil {
		ps.log.Error("[ ProductsStore.AddProduct ] Error inserting categories", slog.String("error", err.Error()))
		return err
	}

	insertImagesQuery := `
		INSERT INTO product_images (product_id, image_url)
		SELECT $1, unnest($2::text[]);`

	imageUrls := make([]string, len(card.Images))
	for i, img := range card.Images {
		imageUrls[i] = img.Url
	}

	_, err = tx.Exec(ctx, insertImagesQuery, productID, imageUrls)
	if err != nil {
		ps.log.Error("[ ProductsStore.AddProduct ] Error inserting images", slog.String("error", err.Error()))
		return err
	}

	// Insert options
	insertOptionsQuery := `
		INSERT INTO product_options
		    (product_id,  values)
		SELECT $1, $2;
	`

	options := optionsToDto(card.Options)

	optionValue, err := json.Marshal(options)
	if err != nil {
		ps.log.Error("[ ProductsStore.AddProduct ] Error marshalling option", slog.String("error", err.Error()))
		return err
	}

	_, err = tx.Exec(ctx, insertOptionsQuery, productID, optionValue)
	if err != nil {
		ps.log.Error("[ ProductsStore.AddProduct ] Error inserting options", slog.String("error", err.Error()))
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		ps.log.Error("[ ProductsStore.AddProduct ] Error committing transaction", slog.String("error", err.Error()))
		return err
	}

	ps.log.Debug("[ ProductsStore.AddProduct ] ProductBase successfully added")

	return nil
}
