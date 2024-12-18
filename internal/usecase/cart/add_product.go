package cart

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cm *CartManager) AddProduct(ctx context.Context, productID uint32, userID uint32) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	cm.log.Info("[CartManager.AddProduct] Started executing", slog.Any("request-id", requestID))

	product, err := cm.cartRepository.GetCartProduct(ctx, productID, userID)
	if err != nil {
		if errors.Is(err, errs.ProductNotInCart) {
			err = cm.cartRepository.AddProduct(ctx, productID, userID)
			if err != nil {
				cm.log.Error("[CartManager.AddProduct] Error adding product", slog.String("error", err.Error()))

				return err
			}

			return nil
		}

		cm.log.Error("[CartManager.AddProduct] Unexpected error", slog.String("error", err.Error()))

		return err
	}

	if product.IsDeleted {
		err = cm.cartRepository.ChangeCartProductDeletedState(ctx, productID, userID)
		if err != nil {
			return err
		}

		return nil
	}

	return errs.ProductAlreadyInCart
}
