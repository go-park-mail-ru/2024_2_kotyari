package search

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
