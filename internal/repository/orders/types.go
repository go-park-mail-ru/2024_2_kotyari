package rorders

import (
	"github.com/google/uuid"
	"time"
)

type getOrdersRow struct {
	OrderID      uuid.UUID
	OrderDate    time.Time
	DeliveryDate time.Time
	ProductID    uint32
	ProductName  string
	ImageURL     string
	TotalPrice   uint32
	Status       string
}

type getOrderByIdRow struct {
	OrderID   uuid.UUID
	OrderDate time.Time
	Date      time.Time
	ProductID uint32
	Cost      uint32
	Count     uint32
	Weight    float32
	Status    string
	Address   string
	Username  string
	Title     string
	ImageURL  string
}
