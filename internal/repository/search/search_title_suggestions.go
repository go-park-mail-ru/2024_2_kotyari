package search

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/jackc/pgx/v5"
)

func (s *SearchStore) GetSearchTitleSuggestions(ctx context.Context, queryParam string) (model.SearchTitleSuggestions, error) {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return model.SearchTitleSuggestions{}, err
	}

	s.log.Info("[SearchStore.GetSearchTitleSuggestions], request-id: ", slog.Any("request-id", requestID))

	const query = `
		select title from products
		where to_tsvector('russian', title) @@ to_tsquery('russian', $1 || ':*');	
	`

	rows, err := s.db.Query(ctx, query, queryParam)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.log.Error("[SearchStore.GetSearchTitleSuggestions] Err no rows")

			return model.SearchTitleSuggestions{}, errs.NoTitlesToSuggest
		}

		s.log.Error("[SearchStore.GetSearchTitleSuggestions] Unexpected error happened")

		return model.SearchTitleSuggestions{}, err
	}

	titlesSet := make(map[string]bool)

	for rows.Next() {
		var title string

		err = rows.Scan(&title)
		if err != nil {
			s.log.Error("[SearchStore.GetSearchTitleSuggestions] Error parsing rows: ", slog.String("error", err.Error()))

			return model.SearchTitleSuggestions{}, err
		}

		words := strings.Split(title, " ")
		for _, word := range words {
			if containsQueryParam(word, queryParam) {
				titlesSet[word] = true
			}
		}
	}

	if len(titlesSet) == 0 {
		return model.SearchTitleSuggestions{}, errs.NoTitlesToSuggest
	}

	suggestions := make([]string, 0, 6)
	for title := range titlesSet {
		suggestions = append(suggestions, title)
		if len(suggestions) == 6 {
			break
		}
	}

	return model.SearchTitleSuggestions{Titles: suggestions}, nil
}

func containsQueryParam(title, q string) bool {
	title = strings.ToLower(title)
	q = strings.ToLower(q)
	return strings.Contains(title, q)
}
