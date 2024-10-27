package orders

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID       string
	ImageURL string
	Name     string
	Cost     int
	Count    int
	Weight   int
}

type Order struct {
	ID           uuid.UUID
	Recipient    string
	Address      string
	OrderStatus  string
	DeliveryDate time.Time
	OrderDate    time.Time
	Products     []Product
}
