package address

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type AddressResponse struct {
	Address string `json:"address"`
}

type UpdateAddressRequest struct {
	Address string `json:"address"`
}

func (a *UpdateAddressRequest) ToModel() model.Address {
	return model.Address{
		Text: a.Address,
	}
}

func FromModel(address model.Address) AddressResponse {
	return AddressResponse{
		Address: address.Text,
	}
}
