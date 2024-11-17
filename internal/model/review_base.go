package model

import "time"

type Review struct {
	ID        uint32
	Username  string
	AvatarURL string
	Text      string
	Rating    uint8
	IsPrivate bool
	UpdatedAt time.Time
	CreatedAt time.Time
}
