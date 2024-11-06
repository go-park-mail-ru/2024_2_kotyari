package model

import "time"

type Cart struct {
	ID           uint32
	UserID       uint32
	DeliveryDate time.Time
	Products     []CartProduct
}
