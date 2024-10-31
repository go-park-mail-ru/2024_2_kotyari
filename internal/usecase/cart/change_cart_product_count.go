package cartServiceLib

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"log/slog"
)

func (am *CartManager) ChangeCartProductCount(ctx context.Context, productID uint32, count int32) error {
	productCount, err := am.cartRepository.GetCartProductCount(ctx, productID)
	if err != nil {
		am.log.Error("[CartManager.ChangeCartProductCount] Error getting productCount count", slog.String("error", err.Error()))

		return err
	}

	switch {
	case count > 0:
		if int32(productCount)-count >= 0 {
			err = am.cartRepository.ChangeCartProductCount(ctx, productID, count)
			if err != nil {
				am.log.Error("[CartManager.ChangeCartProductCount] Error changing product count", slog.String("error", err.Error()))

				return errs.ProductCountTooLow
			}

		}

		return nil

	case count < 0:
		err = am.cartRepository.ChangeCartProductCount(ctx, productID, count)
		if err != nil {
			am.log.Error("[CartManager.ChangeCartProductCount] Error changing product count", slog.String("error", err.Error()))

			return err
		}

	}

	return nil
}
