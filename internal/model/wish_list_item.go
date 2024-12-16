package model

import "time"

type WishlistItem struct {
	ProductID uint32
	AddedAt   time.Time
}
