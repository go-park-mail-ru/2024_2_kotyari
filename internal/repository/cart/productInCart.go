package cart

import (
	"context"
	"fmt"
)

func (cs *CartsStore) ProductInCart(ctx context.Context, userId uint32, productId uint32) (bool, error) {
	const query = `
        SELECT EXISTS (
            SELECT 1 
            FROM carts 
            WHERE user_id = $1 AND product_id = $2 AND is_deleted = false
        )
    `

	var exists bool

	err := cs.db.QueryRow(ctx, query, userId, productId).
		Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка при проверке наличия продукта в корзине: %w", err)
	}

	return exists, nil
}
