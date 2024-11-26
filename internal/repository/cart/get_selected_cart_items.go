package cart

import (
	"context"
	"log/slog"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cs *CartsStore) GetSelectedCartItems(ctx context.Context, userID uint32) ([]order.ProductOrder, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	cs.log.Info("[CartsStore.GetSelectedCartItems] Started executing", slog.Any("request-id", requestID))

	const query = `
		SELECT p.id, p.title, p.image_url, p.price, c.count, p.weight
		FROM carts AS c
		INNER JOIN products AS p ON c.product_id = p.id
		WHERE c.user_id = $1 AND c.is_selected = true AND c.is_deleted = false;
	`

	rows, err := cs.db.Query(ctx, query, userID)
	if err != nil {
		cs.log.Error("[OrdersRepo.GetSelectedCartItems] Failed to retrieve selected items", slog.String("error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var items []order.ProductOrder
	for rows.Next() {
		var item order.ProductOrder
		if err := rows.Scan(&item.ID, &item.Name, &item.ImageUrl, &item.Cost, &item.Count, &item.Weight); err != nil {
			cs.log.Error("[OrdersRepo.GetSelectedCartItems] Failed to scan item", slog.String("error", err.Error()))
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
