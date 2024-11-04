package morders

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
}

type getOrderByIdRow struct {
	orderID   uuid.UUID
	orderDate time.Time
	date      time.Time
	productID uint32
	cost      uint16
	count     uint16
	weight    uint16
	status    string
	address   string
	username  string
	title     string
	imageURL  string
}
