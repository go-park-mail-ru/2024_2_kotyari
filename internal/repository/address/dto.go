package address

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5/pgtype"
)

type AddressDTO struct {
	City   string
	Street string
	House  string
	Flat   pgtype.Text
}

func (a *AddressDTO) ToModel() model.Address {
	return model.Address{
		City:   a.City,
		Street: a.Street,
		House:  a.House,
		Flat:   a.Flat.String,
	}
}
