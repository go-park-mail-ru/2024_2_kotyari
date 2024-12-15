package category

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

func (s *CategoryTestSuite) TestGetAllCategoriesSuccess() {
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `SELECT name, link_to, picture
                   FROM categories
                   WHERE active=true`

	rows := pgxmock.NewRows([]string{"name", "link_to", "picture"}).
		AddRow("Category 1", "link_1", "picture_1").
		AddRow("Category 2", "link_2", "picture_2")

	s.dbMock.ExpectQuery(query).WillReturnRows(rows)

	categories, err := s.store.GetAllCategories(s.ctx)
	require.NoError(s.T(), err)
	require.Len(s.T(), categories, 2)

	require.Equal(s.T(), "Category 1", categories[0].Name)
	require.Equal(s.T(), "link_1", categories[0].LinkTo)
	require.Equal(s.T(), "picture_1", categories[0].Picture)
}

func (s *CategoryTestSuite) TestGetAllCategoriesQueryError() {
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `SELECT name, link_to, picture
                   FROM categories
                   WHERE active=true`

	s.dbMock.ExpectQuery(query).WillReturnError(errors.New("database error"))

	categories, err := s.store.GetAllCategories(s.ctx)
	require.Error(s.T(), err)
	require.Nil(s.T(), categories)
}

func (s *CategoryTestSuite) TestGetAllCategoriesNoRows() {
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `SELECT name, link_to, picture
                   FROM categories
                   WHERE active=true`

	rows := pgxmock.NewRows([]string{"name", "link_to", "picture"})

	s.dbMock.ExpectQuery(query).WillReturnRows(rows)

	categories, err := s.store.GetAllCategories(s.ctx)
	require.ErrorIs(s.T(), err, errs.CategoriesDoesNotExits)
	require.Nil(s.T(), categories)
}

func (s *CategoryTestSuite) TestGetAllCategoriesScanError() {
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `SELECT name, link_to, picture
                   FROM categories
                   WHERE active=true`

	rows := pgxmock.NewRows([]string{"name", "link_to", "picture"}).
		AddRow("Category 1", nil, "picture_1").RowError(0, fmt.Errorf("row scan error"))

	s.dbMock.ExpectQuery(query).WillReturnRows(rows)

	categories, err := s.store.GetAllCategories(s.ctx)
	require.Error(s.T(), err)
	require.Nil(s.T(), categories)
}
