package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cm *CartManager) ChangeCartProductSelectedState(ctx context.Context, productID uint32, userID uint32, isSelected bool) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	cm.log.Info("[CartManager.ChangeCartProductSelectedState] Started executing", slog.Any("request-id", requestID))

	product, err := cm.cartRepository.GetCartProduct(ctx, productID, userID)
	if err != nil {
		cm.log.Error("[CartManager.ChangeCartProductSelectedState] Error getting product", slog.String("error", err.Error()))

		return err
	}

	if product.IsDeleted {
		return errs.ProductNotInCart
	}

	err = cm.cartRepository.ChangeCartProductSelectedState(ctx, productID, userID, isSelected)
	if err != nil {
		cm.log.Error("[CartManager.ChangeCartProductSelectedState] Error changing product selected state", slog.String("error", err.Error()))

		return err
	}

	return nil
}
