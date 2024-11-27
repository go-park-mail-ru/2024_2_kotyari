package search

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (s *SearchHandler) GetSearchTitleSuggestions(w http.ResponseWriter, r *http.Request) {
	requestID, err := utils.GetContextRequestID(r.Context())
	if err != nil {
		s.log.Error("[SearchHandler.GetSearchTitleSuggestions] No request ID")
		utils.WriteErrorJSONByError(w, err, s.errResolver)

		return
	}

	s.log.Info("[SearchHandler.GetSearchTitleSuggestions] Started executing", slog.Any("request-id", requestID))

	query := utils.GetSearchQuery(r)

	titles, err := s.searchRepository.GetSearchTitleSuggestions(r.Context(), query)
	if err != nil {
		err, code := s.errResolver.Get(err)
		utils.WriteErrorJSON(w, code, err)

		return
	}

	var resp GetSuggestionsResponse

	suggestions := make([]Suggestion, 0, len(titles.Titles))
	for _, title := range titles.Titles {
		suggestions = append(suggestions, suggestionFromTitle(title))
	}

	resp = getSuggestionsFromSuggestion(suggestions)
	utils.WriteJSON(w, http.StatusOK, resp)
}
