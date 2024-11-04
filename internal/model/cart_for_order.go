package model

type CartForOrder struct {
	TotalItems             uint16
	TotalWeight            float32
	FinalPrice             float32
	DeliveryDates          []DeliveryDateForOrder
	UserName               string
	PreferredPaymentMethod string
}
