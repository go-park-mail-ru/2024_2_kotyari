package category

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
)

type CategoryTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *CategoriesStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *CategoryTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)
	s.logger = logger.InitLogger()
	s.store = &CategoriesStore{
		db:  s.dbMock,
		log: s.logger,
	}
	requestID := "test-request-id"
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)
}

func (s *CategoryTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *CategoryTestSuite) TestGetProductsByCategoryLinkSuccess() {
	categoryLink := "electronics"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `SELECT p\.id, p\.title, p\.price, p\.original_price, 
                     p\.discount, p\.image_url, p\.description, p\.rating, p\.type, p\.tags
              FROM products p
              JOIN product_categories pc ON p\.id = pc\.product_id
              JOIN categories c ON pc\.category_id = c\.id
              WHERE p\.active = true AND p\.count > 0 AND c\.link_to = \$1
              ORDER BY p\.rating asc;`

	rows := pgxmock.NewRows([]string{"id", "title", "price", "original_price", "discount", "image_url", "description", "rating", "type", "tags"}).
		AddRow(uint32(2), "Product 2", uint32(200), uint32(220), uint32(10), "image_url_2", "description_2", float32(4.7), "Футболка", []string{"черная"}).
		AddRow(uint32(1), "Product 1", uint32(1000), uint32(1200), uint32(20), "image_url_1", "description_1", float32(4.5), "Футболка", []string{"черная"})

	s.dbMock.ExpectQuery(query).WithArgs(categoryLink).WillReturnRows(rows)

	products, err := s.store.GetProductsByCategoryLink(s.ctx, categoryLink, "rating", "asc")
	require.NoError(s.T(), err)
	require.Len(s.T(), products, 2)

	require.Equal(s.T(), uint32(2), products[0].ID)
	require.Equal(s.T(), "Product 2", products[0].Title)
}

func (s *CategoryTestSuite) TestGetProductsByCategoryLinkPriceSuccess() {
	categoryLink := "electronics"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `SELECT p\.id, p\.title, p\.price, p\.original_price, 
                     p\.discount, p\.image_url, p\.description, p\.rating, p\.type, p\.tags
              FROM products p
              JOIN product_categories pc ON p\.id = pc\.product_id
              JOIN categories c ON pc\.category_id = c\.id
              WHERE p\.active = true AND p\.count > 0 AND c\.link_to = \$1
              ORDER BY p\.price asc;`

	rows := pgxmock.NewRows([]string{"id", "title", "price", "original_price", "discount", "image_url", "description", "rating", "type", "tags"}).
		AddRow(uint32(2), "Product 2", uint32(200), uint32(220), uint32(10), "image_url_2", "description_2", float32(4.7), "Футболка", []string{"черная"}).
		AddRow(uint32(1), "Product 1", uint32(1000), uint32(1200), uint32(20), "image_url_1", "description_1", float32(4.5), "Футболка", []string{"черная"})

	s.dbMock.ExpectQuery(query).WithArgs(categoryLink).WillReturnRows(rows)

	products, err := s.store.GetProductsByCategoryLink(s.ctx, categoryLink, "price", "asc")
	require.NoError(s.T(), err)
	require.Len(s.T(), products, 2)

	require.Equal(s.T(), uint32(2), products[0].ID)
	require.Equal(s.T(), "Product 2", products[0].Title)
}

func (s *CategoryTestSuite) TestGetProductsByCategoryLinkNoProducts() {
	categoryLink := "empty-category"

	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `SELECT p\.id, p\.title, p\.price, p\.original_price, 
                     p\.discount, p\.image_url, p\.description, p\.rating, p\.type, p\.tags
              FROM products p
              JOIN product_categories pc ON p\.id = pc\.product_id
              JOIN categories c ON pc\.category_id = c\.id
              WHERE p\.active = true AND p\.count > 0 AND c\.link_to = \$1
              ORDER BY p\.created_at asc;`

	s.dbMock.ExpectQuery(query).WithArgs(categoryLink).WillReturnRows(pgxmock.NewRows(nil))

	products, err := s.store.GetProductsByCategoryLink(s.ctx, categoryLink, "created_at", "asc")
	require.ErrorIs(s.T(), err, errs.ProductsDoesNotExists)
	require.Nil(s.T(), products)
}

func (s *CategoryTestSuite) TestGetProductsByCategoryLinkQueryError() {
	categoryLink := "electronics"

	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `SELECT p\.id, p\.title, p\.price, p\.original_price, 
                     p\.discount, p\.image_url, p\.description, p\.rating, p\.type, p\.tags
              FROM products p
              JOIN product_categories pc ON p\.id = pc\.product_id
              JOIN categories c ON pc\.category_id = c\.id
              WHERE p\.active = true AND p\.count > 0 AND c\.link_to = \$1
              ORDER BY p\.created_at asc;`

	s.dbMock.ExpectQuery(query).WithArgs(categoryLink).WillReturnError(errors.New("query error"))

	products, err := s.store.GetProductsByCategoryLink(s.ctx, categoryLink, "created_at", "asc")
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "query error")
	require.Nil(s.T(), products)
}

func (s *CategoryTestSuite) TestGetProductsByCategoryLinkRowScanError() {
	categoryLink := "electronics"

	query := `SELECT p\.id, p\.title, p\.price, p\.original_price, 
                     p\.discount, p\.image_url, p\.description, p\.rating, p\.type, p\.tags
              FROM products p
              JOIN product_categories pc ON p\.id = pc\.product_id
              JOIN categories c ON pc\.category_id = c\.id
              WHERE p\.active = true AND p\.count > 0 AND c\.link_to = \$1
              ORDER BY p\.created_at asc;`

	rows := pgxmock.NewRows([]string{"id", "title", "price", "original_price", "discount", "image_url", "description", "rating"}).
		AddRow(uint32(1), nil, uint32(1000), uint32(1200), uint32(20), "image_url_1", "description_1", float32(4.5)).RowError(0, fmt.Errorf("row scan error"))

	s.dbMock.ExpectQuery(query).WithArgs(categoryLink).WillReturnRows(rows)

	products, err := s.store.GetProductsByCategoryLink(s.ctx, categoryLink, "created_at", "asc")
	require.Error(s.T(), err)
	require.Nil(s.T(), products)
}

func TestCategoryTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryTestSuite))
}
