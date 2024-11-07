package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductOrder struct {
	ID           uint32
	ProductID    uint32
	OptionID     *uint32
	Count        uint32
	Weight       float32
	Cost         uint32
	OrderID      uuid.UUID
	DeliveryDate time.Time
	Name         string
	ImageUrl     string
}
