package rorders

import (
	"context"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

// Заглушка для получения товаров из корзины
func (r *OrdersRepo) GetCartItems(ctx context.Context, userID uint32) ([]order.ProductOrder, error) {
	return []order.ProductOrder{
		{
			ID:       5,
			Name:     "Product 1",
			ImageUrl: "https://example.com/image1.png",
			Cost:     1500,
			Count:    2,
			Weight:   500,
		},
		{
			ID:       6,
			Name:     "Product 2",
			ImageUrl: "https://example.com/image2.png",
			Cost:     2500,
			Count:    1,
			Weight:   300,
		},
	}, nil
}
