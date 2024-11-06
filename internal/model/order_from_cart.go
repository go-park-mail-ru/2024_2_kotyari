package model

import (
	"github.com/google/uuid"
	"time"
)

type OrderFromCart struct {
	OrderID      uuid.UUID
	UserID       uint32
	Address      string
	TotalPrice   uint32
	DeliveryDate time.Time
	CreatedAt    time.Time
	Products     []ProductOrder
}
