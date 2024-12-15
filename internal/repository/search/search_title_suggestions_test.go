package search

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func (s *SearchTestSuite) TestGetSearchTitleSuggestionsSuccess() {
	searchQuery := "phone"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select title from products where to_tsvector\('russian', title\) @@ to_tsquery\('russian', \$1 \|\| ':\*'\);`

	rows := pgxmock.NewRows([]string{"title"}).
		AddRow("Smartphone Samsung").
		AddRow("iPhone 12").
		AddRow("Phone")

	s.dbMock.ExpectQuery(query).WithArgs(searchQuery).WillReturnRows(rows)

	suggestions, err := s.store.GetSearchTitleSuggestions(s.ctx, searchQuery)
	require.NoError(s.T(), err)
	require.Len(s.T(), suggestions.Titles, 3)
	require.Contains(s.T(), suggestions.Titles, "Smartphone")
	require.Contains(s.T(), suggestions.Titles, "Phone")
}

func (s *SearchTestSuite) TestGetSearchTitleSuggestionsNoResults() {
	searchQuery := "nonexistent"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select title from products where to_tsvector\('russian', title\) @@ to_tsquery\('russian', \$1 \|\| ':\*'\);`

	s.dbMock.ExpectQuery(query).WithArgs(searchQuery).WillReturnRows(pgxmock.NewRows(nil))

	suggestions, err := s.store.GetSearchTitleSuggestions(s.ctx, searchQuery)
	require.ErrorIs(s.T(), err, errs.NoTitlesToSuggest)
	require.Empty(s.T(), suggestions.Titles)
}

func (s *SearchTestSuite) TestGetSearchTitleSuggestionsQueryError() {
	searchQuery := "query-error"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select title from products where to_tsvector\('russian', title\) @@ to_tsquery\('russian', \$1 \|\| ':\*'\);`

	s.dbMock.ExpectQuery(query).WithArgs(searchQuery).WillReturnError(errors.New("query error"))

	suggestions, err := s.store.GetSearchTitleSuggestions(s.ctx, searchQuery)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "query error")
	require.Empty(s.T(), suggestions.Titles)
}

func (s *SearchTestSuite) TestGetSearchTitleSuggestionsRowScanError() {
	searchQuery := "row-scan-error"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select title from products where to_tsvector\('russian', title\) @@ to_tsquery\('russian', \$1 \|\| ':\*'\);`

	rows := pgxmock.NewRows([]string{"title"}).
		AddRow(nil).RowError(0, errors.New("row scan error"))

	s.dbMock.ExpectQuery(query).WithArgs(searchQuery).WillReturnRows(rows)

	suggestions, err := s.store.GetSearchTitleSuggestions(s.ctx, searchQuery)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "row scan error")
	require.Empty(s.T(), suggestions.Titles)
}

func TestContainsQueryParam(t *testing.T) {
	require.True(t, containsQueryParam("Samsung", "sam"))
	require.False(t, containsQueryParam("Samsung", "apple"))
	require.True(t, containsQueryParam("Samsung Galaxy", "gal"))
	require.False(t, containsQueryParam("", "any"))
}
