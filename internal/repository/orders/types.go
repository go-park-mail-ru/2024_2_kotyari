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
	TotalPrice   uint16
	Status       string
}

type getOrderByIdRow struct {
	OrderID   uuid.UUID
	OrderDate time.Time
	Date      time.Time
	ProductID uint32
	Cost      uint16
	Count     uint16
	Weight    uint16
	Status    string
	Address   string
	Username  string
	Title     string
	ImageURL  string
}
