package model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID           uuid.UUID
	DeliveryDate time.Time
	OrderDate    time.Time
	TotalPrice   uint16
	Status       string
	Recipient    string
	Address      string
	Products     []ProductOrder
}