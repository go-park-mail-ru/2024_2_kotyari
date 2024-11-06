package model

import "time"

type DeliveryDateForOrder struct {
	Date   time.Time
	Weight float32
	Items  []CartProductForOrder
}
