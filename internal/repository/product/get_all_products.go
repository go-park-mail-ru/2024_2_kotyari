package product

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (ps *ProductsStore) GetAllProducts(ctx context.Context) ([]model.ProductCatalog, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	ps.log.Info("[ProductsStore.GetAllProducts] Started executing", slog.Any("request-id", requestID))

	const query = `
		SELECT p.id, p.title, p.price, p.original_price, p.rating,
		       p.discount, p.image_url, p.description
		FROM products p
			where p.active = true and p.count > 0
		ORDER BY p.created_at DESC;
	`
	ps.log.Debug("[ ProductsStore.GetAllProducts ] is running! ]")

	rows, err := ps.db.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ps.log.Info("[ProductsStore.GetAllProducts] нет продуктов в БД")
			return nil, errs.ProductsDoesNotExists
		}

		ps.log.Error("[ProductsStore.GetAllProducts] Query error ", slog.String("error", err.Error()))
	}

	defer rows.Close()

	var products []model.ProductCatalog
	for rows.Next() {
		var p model.ProductCatalog

		err = rows.Scan(
			&p.ID, &p.Title, &p.Price, &p.OriginalPrice, &p.Rating,
			&p.Discount, &p.ImageURL, &p.Description,
		)
		if err != nil {
			ps.log.Error("[ProductsStore.GetAllProducts] rows.Scan error ", slog.String("error", err.Error()))

			return nil, err
		}

		products = append(products, p)
	}

	ps.log.Info("[ProductsStore.GetAllProducts] success get products", slog.Int("count", len(products)))

	return products, nil
}
