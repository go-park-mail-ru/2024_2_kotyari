package utils

type MessageType string

const (
	AddPromo    MessageType = "add_promo"
	DeletePromo MessageType = "remove_promo"
)

type PromoMessage struct {
	UserID  uint32 `json:"user_id"`
	PromoID uint32 `json:"promo_id"`
}
