package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cm *CartManager) ChangeCartProductCount(ctx context.Context, productID uint32, count int32, userID uint32) (uint32, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return 0, err
	}

	cm.log.Info("[CartManager.ChangeCartProductCount] Started executing", slog.Any("request-id", requestID))

	product, err := cm.cartRepository.GetCartProduct(ctx, productID, userID)
	if err != nil {
		cm.log.Error("[CartManager.ChangeCartProductCount] Error getting product count", slog.String("error", err.Error()))

		return 0, err
	}

	if product.IsDeleted {
		return 0, errs.ProductNotInCart
	}

	productCount, err := cm.productCountGetter.GetProductCount(ctx, productID)
	if err != nil {
		cm.log.Error("[CartManager.ChangeCartProductCount] Error getting productCount count", slog.String("error", err.Error()))

		return 0, err
	}

	err = cm.validateCartProductCount(ctx, count, productCount, product.Count, productID, userID)
	if err != nil {
		cm.log.Error("[CartManager.ChangeCartProductCount] Error changing product count", slog.String("error", err.Error()))

		return 0, err
	}

	cartProductCount, err := cm.cartRepository.GetCartProductCount(ctx, userID, productID)
	if err != nil {
		cm.log.Error("[CartManager.ChangeCartProductCount] Error getting cart product count",
			slog.String("error", err.Error()))

		return 0, err
	}

	return cartProductCount, nil
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
			err := cm.cartRepository.RemoveCartProduct(ctx, productID, count, userID)
			if err != nil {
				cm.log.Error("[CartManager.validateCartProductCount] Error removing cart", slog.String("error", err.Error()))

				return err
			}

			return nil
		}
	}

	return errs.BadRequest
}
