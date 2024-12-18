package notifications

import notifications "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/notifications/gen"

type OrderUpdateResponse struct {
	Updates []OrderUpdate `json:"orders_updates"`
}

type OrderUpdate struct {
	OrderID     string `json:"order_id"`
	OrderStatus string `json:"new_status"`
}

func orderUpdateFromGrpc(orderUpdate *notifications.OrderUpdateMessage) OrderUpdate {
	return OrderUpdate{
		OrderID:     orderUpdate.GetOrderId(),
		OrderStatus: orderUpdate.GetNewStatus(),
	}
}
