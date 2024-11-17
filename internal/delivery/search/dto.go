package search

import "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"

type GetSuggestionsResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
}

func getSuggestionsFromSuggestion(suggestions []Suggestion) GetSuggestionsResponse {
	return GetSuggestionsResponse{
		Suggestions: suggestions,
	}
}

type Suggestion struct {
	Title string `json:"title"`
}

func suggestionFromTitle(title string) Suggestion {
	return Suggestion{Title: title}
}

type dtoProductCatalog struct {
	dtoProduct
	ImageURL string `json:"image_url"`
}

type dtoProduct struct {
	ID            uint32  `json:"id"`
	Description   string  `json:"description"`
	Count         uint32  `json:"count"`
	Title         string  `json:"title"`
	Price         uint32  `json:"price"`
	OriginalPrice uint32  `json:"original_price"`
	Discount      uint32  `json:"discount"`
	Rating        float32 `json:"rating"`
}

func newDTOProductCatalogFromModel(pc model.ProductCatalog) dtoProductCatalog {
	return dtoProductCatalog{
		dtoProduct: dtoProduct{
			ID:            pc.ID,
			Description:   pc.Description,
			Count:         pc.Count,
			Title:         pc.Title,
			Price:         pc.Price,
			OriginalPrice: pc.OriginalPrice,
			Discount:      pc.Discount,
			Rating:        pc.Rating,
		},
		ImageURL: pc.ImageURL,
	}
}
