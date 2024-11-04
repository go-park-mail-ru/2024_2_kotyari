package cart

import "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"

var (
	availablePaymentMethods = []string{"Картой", "Наличными"}
	currency                = "₽"
)

func getPaymentIcon(method string) string {
	switch method {
	case "Картой":
		return "credit_card"
	case "Наличными":
		return "payments"
	default:
		return "payment"
	}
}

func cartForOrderResponseFromModel(cart model.CartForOrder) orderData {
	deliveryDates := make([]deliveryDateInfo, 0, len(cart.DeliveryDates))

	for _, deliveryDate := range cart.DeliveryDates {
		items := make([]productResponse, 0, len(deliveryDate.Items))
		for _, item := range deliveryDate.Items {
			items = append(items, productResponse{
				Title:    item.Title,
				Price:    item.Price,
				Quantity: item.Quantity,
				Image:    item.Image,
				Weight:   item.Weight,
				URL:      item.URL,
			})
		}
		deliveryDates = append(deliveryDates, deliveryDateInfo{
			Date:   deliveryDate.Date,
			Weight: deliveryDate.Weight,
			Items:  items,
		})
	}

	paymentMethods := make([]paymentMethod, len(availablePaymentMethods))
	for i, method := range availablePaymentMethods {
		paymentMethods[i] = paymentMethod{
			Method:     method,
			Icon:       getPaymentIcon(method),
			IsSelected: method == cart.PreferredPaymentMethod,
		}
	}

	return orderData{
		TotalItems:     cart.TotalItems,
		TotalWeight:    cart.TotalWeight,
		FinalPrice:     cart.FinalPrice,
		Currency:       currency,
		PaymentMethods: paymentMethods,
		Recipient: recipientInfo{
			Address:       "г. Москва, 2-я Бауманская ул., 5",
			RecipientName: cart.UserName,
		},
		DeliveryDates: deliveryDates,
	}
}
