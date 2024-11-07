package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID
	DeliveryDate time.Time
	OrderDate    time.Time
	TotalPrice   uint32
	Status       string
	Recipient    string
	Address      string
	Products     []ProductOrder
}
