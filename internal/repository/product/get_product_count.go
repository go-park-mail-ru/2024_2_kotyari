package product

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (ps *ProductsStore) GetProductCount(ctx context.Context, productID uint32) (uint32, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return 0, err
	}

	ps.log.Info("[ProductsStore.GetProductCount] Started executing", slog.Any("request-id", requestID))

	const query = `
		select count from products 
		where id = $1;
	`

	var count uint32

	err = ps.db.QueryRow(ctx, query, productID).Scan(&count)
	if err != nil {
		ps.log.Error("[CartStore.GetProductCount] Error performing query: ", slog.String("error", err.Error()))

		return 0, err
	}

	return count, nil
}
