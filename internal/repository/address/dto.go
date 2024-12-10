package address

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5/pgtype"
)

type AddressDTO struct {
	Address pgtype.Text
}

func (a *AddressDTO) ToModel() model.Address {
	return model.Address{
		Text: a.Address.String,
	}
}
