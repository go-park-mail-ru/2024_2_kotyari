package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

func (cm *CartManager) ChangeCartProductCount(ctx context.Context, productID uint32, count int32, userID uint32) error {
	product, err := cm.cartRepository.GetCartProduct(ctx, productID, userID)
	if err != nil {
		cm.log.Error("[CartManager.ChangeCartProductCount] Error getting product count", slog.String("error", err.Error()))

		return err
	}

	if product.IsDeleted {
		return errs.ProductNotInCart
	}

	productCount, err := cm.productCountGetter.GetProductCount(ctx, productID)
	if err != nil {
		cm.log.Error("[CartManager.ChangeCartProductCount] Error getting productCount count", slog.String("error", err.Error()))

		return err
	}

	err = cm.validateCartProductCount(ctx, count, productCount, product.Count, productID, userID)
	if err != nil {
		cm.log.Error("[CartManager.ChangeCartProductCount] Error changing product count", slog.String("error", err.Error()))

		return err
	}

	return nil
}

func (cm *CartManager) validateCartProductCount(ctx context.Context, count int32, productCount uint32, cartProductCount uint32, productID uint32, userID uint32) error {
	switch {
	case count > 0:
		if int32(productCount)-count >= 0 {
			err := cm.cartRepository.ChangeCartProductCount(ctx, productID, count, userID)
			if err != nil {
				cm.log.Error("[CartManager.validateCartProductCount] Error changing productCount count", slog.String("error", err.Error()))

				return err
			}

			return nil
		}

		return errs.ProductCountTooLow

	case count < 0:
		if int32(cartProductCount)+count >= 1 {
			err := cm.cartRepository.ChangeCartProductCount(ctx, productID, count, userID)
			if err != nil {
				cm.log.Error("[CartManager.validateCartProductCount] Error changing productCount count", slog.String("error", err.Error()))

				return err
			}

			return nil
		}

		if int32(cartProductCount)+count == 0 {
			err := cm.cartRepository.RemoveCartProduct(ctx, productID, count)
			if err != nil {
				cm.log.Error("[CartManager.validateCartProductCount] Error removing cart", slog.String("error", err.Error()))

				return err
			}

			return nil
		}
	}

	return errs.BadRequest
}
