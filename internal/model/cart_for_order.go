package model

type CartForOrder struct {
	TotalItems             uint32
	TotalWeight            float32
	FinalPrice             float32
	DeliveryDates          []DeliveryDateForOrder
	UserName               string
	PreferredPaymentMethod string
	Address                AddressInfo
}
