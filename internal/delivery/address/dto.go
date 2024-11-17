package address

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type AddressResponse struct {
	City   string `json:"city"`
	Street string `json:"street"`
	House  string `json:"house"`
	Flat   string `json:"flat"`
}

type UpdateAddressRequest struct {
	City   string `json:"city"`
	Street string `json:"street"`
	House  string `json:"house"`
	Flat   string `json:"flat"`
}

func (a *UpdateAddressRequest) ToModel() model.Address {
	return model.Address{
		City:   a.City,
		Street: a.Street,
		House:  a.House,
		Flat:   &a.Flat,
	}
}

func FromModel(address model.Address) AddressResponse {
	return AddressResponse{
		City:   address.City,
		Street: address.Street,
		House:  address.House,
		Flat:   *address.Flat,
	}
}
