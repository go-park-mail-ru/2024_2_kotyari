package promocodes

import (
	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

func promoCodeToGRPC(promoCode model.PromoCode) *promocodes.PromoCode {
	return &promocodes.PromoCode{
		Id:     promoCode.ID,
		UserId: promoCode.UserID,
		Name:   promoCode.Name,
		Bonus:  promoCode.Bonus,
	}
}
