package orders

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID           uuid.UUID
	Recipient    string
	Address      string
	OrderStatus  string
	DeliveryDate time.Time
	OrderDate    time.Time
	Products     []Product
}
