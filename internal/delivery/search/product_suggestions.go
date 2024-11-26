package search

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (s *SearchHandler) ProductSuggestions(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		s.log.Error("[SearchHandler.ProductSuggestions] No request ID")
		utils.WriteErrorJSONByError(w, err, s.errResolver)

		return
	}

	s.log.Info("[SearchHandler.ProductSuggestions] Started executing", slog.Any("request-id", requestID))

	query := utils.GetSearchQuery(r)

	sortField := r.URL.Query().Get(utils.SearchFieldParam)
	sortOrder := r.URL.Query().Get(utils.SearchOrderParam)

	products, err := s.searchRepository.ProductSuggestion(r.Context(), query, sortField, sortOrder)
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
