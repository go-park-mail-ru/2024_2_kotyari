package model

type CartForOrder struct {
	TotalItems             uint32
	TotalWeight            float32
	FinalPrice             uint32
	DeliveryDates          []DeliveryDateForOrder
	UserName               string
	PreferredPaymentMethod string
	Address                Address
}
