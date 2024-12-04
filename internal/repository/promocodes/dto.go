package promocodes

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"time"
)

type PromoCodesDTO struct {
	ID        uint32    `db:"id"`
	UserID    uint32    `db:"user_id"`
	Name      string    `db:"name"`
	Bonus     uint32    `db:"promo"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (p *PromoCodesDTO) ToModel() model.PromoCode {
	return model.PromoCode{
		ID:     p.ID,
		UserID: p.UserID,
		Name:   p.Name,
		Bonus:  p.Bonus,
	}
}

func PromoCodesToModelSlice(promoCodes []PromoCodesDTO) []model.PromoCode {
	promoCodesModel := make([]model.PromoCode, 0, len(promoCodes))

	for _, promoCode := range promoCodes {
		promoCodesModel = append(promoCodesModel, promoCode.ToModel())
	}

	return promoCodesModel
}
