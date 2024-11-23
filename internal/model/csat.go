package model

import "time"

type CSATType string

const (
	SurveyTypeSite     CSATType = "site"
	SurveyTypePurchase CSATType = "purchase"
)

type CSAT struct {
	ID        uint32
	UserID    uint32
	Rating    uint32
	Type      CSATType
	Text      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
