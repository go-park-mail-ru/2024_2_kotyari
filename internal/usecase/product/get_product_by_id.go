package product

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (p *ProductService) GetProductByID(ctx context.Context, userID uint32, productID uint32) (model.ProductCard, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		p.log.Error("[ProductService.GetProductByID] Failed to get request-id",
			slog.String("error", err.Error()))

		return model.ProductCard{}, err
	}

	p.log.Info("[ProductService.GetProductByID] Started executing", slog.Any("request-id", requestID))

	productCard, err := p.productCardGetter.GetProductByID(ctx, productID)
	if err != nil {
		p.log.Info("[ProductService.GetProductByID] Error getting product card",
			slog.String("error", err.Error()))

		return model.ProductCard{}, err
	}

	if userID == 0 {
		productCard.IsInCart = false
		productCard.Count = 0

		return productCard, nil
	}

	ok, err := p.cartManager.ProductInCart(ctx, userID, productID)
	if err != nil {
		p.log.Error("[ProductService.GetProductByID] Failed to check if product is in cart",
			slog.String("error", err.Error()))

		return model.ProductCard{}, err
	}

	if !ok {
		productCard.IsInCart = false
		productCard.Count = 0

		return productCard, nil
	}

	cartProductCount, err := p.cartManager.GetCartProductCount(ctx, userID, productID)
	if err != nil {
		p.log.Error("[ProductService.GetProductByID] Failed to get product count",
			slog.String("error", err.Error()))

		return model.ProductCard{}, err
	}

	productCard.Count = cartProductCount
	productCard.IsInCart = true

	return productCard, nil
}
