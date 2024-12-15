package utils

import "github.com/google/uuid"

type MessageType string

const (
	AddPromo            MessageType = "add_promo"
	DeletePromo         MessageType = "remove_promo"
	AvailPromoTenID                 = 1
	AvailPromoTwoFiveID             = 2
	PromoTopic                      = "promo-topic"
)

type PromoMessage struct {
	UserID    uint32    `json:"user_id"`
	PromoID   uint32    `json:"promo_id"`
	RequestID uuid.UUID `json:"request_id"`
}
