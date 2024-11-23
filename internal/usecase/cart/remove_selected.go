package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

func (cm *CartManager) RemoveSelected(ctx context.Context, userID uint32) error {
	products, err := cm.cartRepository.GetSelectedCartItems(ctx, userID)
	if err != nil {
		cm.log.Error("[CartManager.RemoveSelected] Error getting selected from cart", slog.String("error", err.Error()))

		return err
	}

	if len(products) == 0 {
		return errs.NoSelectedProducts
	}

	for _, product := range products {
		err := cm.cartRepository.RemoveCartProduct(ctx, product.ID, int32(product.Count), userID)
		if err != nil {
			cm.log.Error("[CartManager.RemoveSelected] Error deleting selected product from cart",
				slog.String("error", err.Error()),
				slog.Uint64("product_id", uint64(product.ProductID)))
			return err
		}
	}

	return nil
}
