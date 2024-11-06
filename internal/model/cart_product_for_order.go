package model

import "time"

type CartProductForOrder struct {
	Title        string
	Price        float32
	Quantity     uint32
	Image        string
	Weight       float32
	URL          string
	DeliveryDate time.Time
}
