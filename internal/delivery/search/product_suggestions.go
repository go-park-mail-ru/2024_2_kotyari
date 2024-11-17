package search

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (s *SearchHandler) ProductSuggestions(w http.ResponseWriter, r *http.Request) {
	query := utils.GetSearchQuery(r)

	products, err := s.searchRepository.ProductSuggestion(r.Context(), query)
	if err != nil {
		err, code := s.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	dtoProducts := make([]dtoProductCatalog, len(products))
	for i, product := range products {
		dtoProducts[i] = newDTOProductCatalogFromModel(product)
	}

	utils.WriteJSON(w, http.StatusOK, dtoProducts)
}
