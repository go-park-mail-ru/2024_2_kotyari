package promocodes

import (
	promocodes "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/promocodes/gen"
)

type UserPromoCodesResponseDTO struct {
	PromoCodes []PromoCodesResponseDTO `json:"promocodes"`
}

type PromoCodesResponseDTO struct {
	Name  string `json:"name"`
	Bonus uint32 `json:"bonus"`
}

func promoCodeFromGrpc(promocode *promocodes.PromoCode) PromoCodesResponseDTO {
	return PromoCodesResponseDTO{
		Name:  promocode.GetName(),
		Bonus: promocode.GetBonus(),
	}
}

func promoCodesFromDTOSlice(promocodes []PromoCodesResponseDTO) UserPromoCodesResponseDTO {
	return UserPromoCodesResponseDTO{
		PromoCodes: promocodes,
	}
}
