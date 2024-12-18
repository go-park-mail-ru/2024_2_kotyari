package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cm *CartManager) RemoveProduct(ctx context.Context, productID uint32, userID uint32) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	cm.log.Info("[CartManager.RemoveProduct] Started executing", slog.Any("request-id", requestID))

	product, err := cm.cartRepository.GetCartProduct(ctx, productID, userID)
	if err != nil {
		cm.log.Error("[CartManager.RemoveProduct] Error retrieving product", slog.String("error", err.Error()))

		return err
	}

	err = cm.cartRepository.RemoveCartProduct(ctx, productID, -int32(product.Count), userID)
	if err != nil {
		cm.log.Error("[CartManager.RemoveProduct] Error removing product", slog.String("error", err.Error()))

		return err
	}

	return nil
}
