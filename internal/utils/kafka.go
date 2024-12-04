package utils

type MessageType string

const (
	AddPromo            MessageType = "add_promo"
	DeletePromo         MessageType = "remove_promo"
	AvailPromoTenID                 = 1
	AvailPromoTwoFiveID             = 2
)

type PromoMessage struct {
	UserID  uint32 `json:"user_id"`
	PromoID uint32 `json:"promo_id"`
}
