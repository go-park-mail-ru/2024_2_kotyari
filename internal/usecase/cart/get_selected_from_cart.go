package cart

import (
	"context"
	"errors"
	"log/slog"
	"math"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cm *CartManager) GetSelectedFromCart(ctx context.Context, userID uint32, promoName string) (model.CartForOrder, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.CartForOrder{}, err
	}

	cm.log.Info("[CartManager.GetSelectedFromCart] Started executing", slog.Any("request-id", requestID))

	products, err := cm.cartRepository.GetSelectedFromCart(ctx, userID)
	if err != nil {
		return model.CartForOrder{}, err
	}

	cart := model.CartForOrder{}

	deliveryDatesMap := make(map[time.Time]*model.DeliveryDateForOrder)

	for _, product := range products.Items {
		totalProductWeight := product.Weight * float32(product.Quantity)
		cart.TotalItems += product.Quantity
		cart.TotalWeight += totalProductWeight
		cart.FinalPrice += product.Price * product.Quantity

		if _, exists := deliveryDatesMap[product.DeliveryDate]; !exists {
			deliveryDatesMap[product.DeliveryDate] = &model.DeliveryDateForOrder{
				Date:   product.DeliveryDate,
				Weight: 0,
				Items:  []model.CartProductForOrder{},
			}
		}
		deliveryDatesMap[product.DeliveryDate].Weight += float32(math.Round(float64(totalProductWeight)*100) / 100)
		deliveryDatesMap[product.DeliveryDate].Items = append(deliveryDatesMap[product.DeliveryDate].Items, product)
	}

	cart.DeliveryDates = make([]model.DeliveryDateForOrder, 0, len(deliveryDatesMap))
	for _, dateInfo := range deliveryDatesMap {
		cart.DeliveryDates = append(cart.DeliveryDates, *dateInfo)
	}

	cart.UserName = products.UserName
	cart.PreferredPaymentMethod = products.PreferredPaymentMethod
	cart.Address = products.Address

	if promoName != "" {
		promoCode, err := cm.promoCodeGetter.GetPromoCode(ctx, userID, promoName)
		if err != nil {
			if errors.Is(err, errs.NoPromoCode) {
				cm.log.Error("[CartManager.GetSelectedFromCart] no promo code")

				return cart, err
			}

			cm.log.Error("[CartManager.GetSelectedFromCart] Unexpected error",
				slog.String("error", err.Error()))

			return cart, nil
		}

		discountAmount := (cart.FinalPrice * promoCode.Bonus) / 100
		cart.FinalPrice -= discountAmount
	}

	return cart, nil
}
