package model

import (
	"github.com/google/uuid"
	"time"
)

type ProductOrder struct {
	ID           uint32
	ProductID    uint32
	OptionID     *uint32
	Count        uint16
	Weight       uint16
	Cost         uint16
	OrderID      uuid.UUID
	DeliveryDate time.Time
	Name         string
	ImageUrl     string
}
