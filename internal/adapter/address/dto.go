package address

import "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"

type AddressAdatperResponseDTO struct {
	Results []struct {
		Title struct {
			Text string `json:"text"`
			Hl   []struct {
				Begin int `json:"begin"`
				End   int `json:"end"`
			} `json:"hl"`
		} `json:"title"`
		Subtitle struct {
			Text string `json:"text"`
		} `json:"subtitle"`
		Tags     []string `json:"tags"`
		Distance struct {
			Text  string  `json:"text"`
			Value float64 `json:"value"`
		} `json:"distance"`
		Address struct {
			FormattedAddress string `json:"formatted_address"`
			Component        []struct {
				Name string   `json:"name"`
				Kind []string `json:"kind"`
			} `json:"component"`
		} `json:"address"`
		URI string `json:"uri"`
	} `json:"results"`
}

func (a AddressAdatperResponseDTO) ToModelSlice() []model.Addresses {
	var addressSlice []model.Addresses

	for _, address := range a.Results {
		addressSlice = append(addressSlice, model.Addresses{Address: address.Title.Text})
	}

	return addressSlice
}

func singleResultToModel(result string) model.Addresses {
	return model.Addresses{
		Address: result,
	}
}
