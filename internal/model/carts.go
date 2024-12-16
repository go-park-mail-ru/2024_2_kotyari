package model

import "time"

type Cart struct {
	ID           uint32
	UserID       uint32
	TotalWeight  float32
	DeliveryDate time.Time
	Products     []CartProduct
}
