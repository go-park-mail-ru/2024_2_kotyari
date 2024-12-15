package search

import (
	"context"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
)

type SearchTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *SearchStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *SearchTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &SearchStore{db: s.dbMock, log: s.logger}
}

func (s *SearchTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *SearchTestSuite) TestProductSuggestionSuccess() {
	searchQuery := "phone"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `SELECT p\.id, p\.title, p\.price, p\.original_price, p\.rating, 
		       p\.discount, p\.image_url, p\.description
		FROM products p
		    where p\.active = true and p\.count > 0 and to_tsvector\('russian', title\) @@ to_tsquery\('russian', \$1 \|\| ':\*'\)
		ORDER BY p\.created_at asc;`

	rows := pgxmock.NewRows([]string{"id", "title", "price", "original_price", "rating", "discount", "image_url", "description"}).
		AddRow(uint32(1), "Product 1", uint32(100), uint32(120), float32(4.5), uint32(20), "image_url_1", "description_1").
		AddRow(uint32(2), "Product 2", uint32(200), uint32(220), float32(4.7), uint32(10), "image_url_2", "description_2")

	s.dbMock.ExpectQuery(query).WithArgs(searchQuery).WillReturnRows(rows)

	products, err := s.store.ProductSuggestion(s.ctx, searchQuery, "created_at", "asc")
	require.NoError(s.T(), err)
	require.Len(s.T(), products, 2)

	require.Equal(s.T(), uint32(1), products[0].ID)
	require.Equal(s.T(), "Product 1", products[0].Title)
}

func (s *SearchTestSuite) TestProductSuggestionNoProducts() {
	searchQuery := "empty-query"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `SELECT p\.id, p\.title, p\.price, p\.original_price, p\.rating, 
		       p\.discount, p\.image_url, p\.description
		FROM products p
		    where p\.active = true and p\.count > 0 and to_tsvector\('russian', title\) @@ to_tsquery\('russian', \$1 \|\| ':\*'\)
		ORDER BY p\.created_at asc;`

	s.dbMock.ExpectQuery(query).WithArgs(searchQuery).WillReturnRows(pgxmock.NewRows(nil))

	products, _ := s.store.ProductSuggestion(s.ctx, searchQuery, "created_at", "asc")
	require.Nil(s.T(), products)
}

//func (s *SearchTestSuite) TestProductSuggestionQueryError() {
//	searchQuery := "query-error"
//	requestID := uuid.New()
//	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)
//
//	query := `SELECT p\.id, p\.title, p\.price, p\.original_price, p\.rating,
//               p\.discount, p\.image_url, p\.description
//        FROM products p
//        WHERE p\.active = true AND p\.count > 0 AND to_tsvector\('russian', title\) @@ to_tsquery\('russian', \$1 \|\| ':\*'\)
//        ORDER BY p\.created_at ASC;`
//
//	s.dbMock.ExpectQuery(query).WithArgs(searchQuery).WillReturnError(errors.New("query error"))
//
//	products, err := s.store.ProductSuggestion(s.ctx, searchQuery, "created_at", "asc")
//
//	require.Error(s.T(), err)
//	require.Contains(s.T(), err.Error(), "query error")
//	require.Nil(s.T(), products)
//}
//
//func (s *SearchTestSuite) TestProductSuggestionRowScanError() {
//	searchQuery := "row-scan-error"
//
//	query := `SELECT p\.id, p\.title, p\.price, p\.original_price, p\.rating,
//		       p\.discount, p\.image_url, p\.description
//		FROM products p
//		    where p\.active = true and p\.count > 0 and to_tsvector\('russian', title\) @@ to_tsquery\('russian', \$1 \|\| ':\*'\)
//		ORDER BY p\.created_at asc;`
//
//	rows := pgxmock.NewRows([]string{"id", "title", "price", "original_price", "rating", "discount", "image_url", "description"}).
//		AddRow(uint32(1), nil, uint32(100), uint32(120), float32(4.5), uint32(20), "image_url_1", "description_1").RowError(0, fmt.Errorf("row scan error"))
//
//	s.dbMock.ExpectQuery(query).WithArgs(searchQuery).WillReturnRows(rows)
//
//	products, err := s.store.ProductSuggestion(s.ctx, searchQuery, "created_at", "asc")
//	require.Error(s.T(), err)
//	require.Nil(s.T(), products)
//}

func TestSearchTestSuite(t *testing.T) {
	suite.Run(t, new(SearchTestSuite))
}
