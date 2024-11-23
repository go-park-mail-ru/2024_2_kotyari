package cart

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func (cm *CartManager) GetSelectedFromCart(ctx context.Context, userID uint32) (model.CartForOrder, error) {
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
		deliveryDatesMap[product.DeliveryDate].Weight += totalProductWeight
		deliveryDatesMap[product.DeliveryDate].Items = append(deliveryDatesMap[product.DeliveryDate].Items, product)
	}

	cart.DeliveryDates = make([]model.DeliveryDateForOrder, 0, len(deliveryDatesMap))
	for _, dateInfo := range deliveryDatesMap {
		cart.DeliveryDates = append(cart.DeliveryDates, *dateInfo)
	}

	cart.UserName = products.UserName
	cart.PreferredPaymentMethod = products.PreferredPaymentMethod
	cart.Address = products.Address

	return cart, nil
}
