package address

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type AddressDTO struct {
	ID        uint32 `json:"id"`
	City      string `json:"city"`
	Street    string `json:"street"`
	House     string `json:"house"`
	Flat      string `json:"flat"`
	ProfileID uint32 `json:"profile_id"`
}

type AddressRequest struct {
	City   string `json:"city"`
	Street string `json:"street"`
	House  string `json:"house"`
	Flat   string `json:"flat"`
}

func (a *AddressDTO) ToModel() model.Address {
	return model.Address{
		Id:     a.ID,
		City:   a.City,
		Street: a.Street,
		House:  a.House,
		Flat:   a.Flat,
	}
}

func FromModel(address model.Address) AddressDTO {
	return AddressDTO{
		ID:     address.Id,
		City:   address.City,
		Street: address.Street,
		House:  address.House,
		Flat:   address.Flat,
	}
}

type AddressDelete struct {
	ProfileId uint32 `json:"profile_id"`
	Id        uint32 `json:"id"`
}
