package model

import (
	"github.com/google/uuid"
	"time"
)

type OrderRow struct {
	OrderID      uuid.UUID
	OrderDate    time.Time
	DeliveryDate time.Time
	ProductID    string
	ImageURL     string
	ProductName  string
}
