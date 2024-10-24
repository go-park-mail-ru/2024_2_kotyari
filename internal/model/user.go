package model

import "time"

type User struct {
	ID           uint32
	Email        string
	Username     string
	City         string
	Age          uint8
	AvatarUrl    string
	Password     string
	Blocked      bool
	BlockedUntil time.Time
}
