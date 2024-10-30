package cart

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cs *CartsStore) ChangeProductQuantity(ctx context.Context, productID uint32, count int32) error {
	userID := utils.GetContextSessionUserID(ctx)

	const query = `
		update products
		set 
	`
}
