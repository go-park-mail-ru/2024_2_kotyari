package address

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type GetAddressResponse struct {
	Address string `json:"address"`
}

type UpdateAddressRequest struct {
	Address string `json:"address"`
}

type AddressesSuggestions struct {
	Addresses []GetAddressResponse `json:"addresses"`
}

func (a *UpdateAddressRequest) ToModel() model.Addresses {
	return model.Addresses{
		Address: a.Address,
	}
}

func addressFromModel(a model.Addresses) GetAddressResponse {
	return GetAddressResponse{
		Address: a.Address,
	}
}

func addressesSuggestionsFromSlice(addresses []GetAddressResponse) AddressesSuggestions {
	return AddressesSuggestions{
		Addresses: addresses,
	}
}
