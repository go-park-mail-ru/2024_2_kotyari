package wish_list_link

import (
	"context"
	"errors"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type WishListLinksRepoTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	repo   *WishListLinksRepo
	ctx    context.Context
}

func (s *WishListLinksRepoTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.repo = &WishListLinksRepo{
		db: s.dbMock,
	}

	s.ctx = context.Background()
}

func (s *WishListLinksRepoTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *WishListLinksRepoTestSuite) TestCreateLinkSuccess() {
	userID := uint32(123)
	link := "http://example.com"

	const query = `(?i)INSERT INTO wish_list_links\s*\(user_id,\s*link\)\s*VALUES\s*\(\$1,\s*\$2\)`

	s.dbMock.ExpectExec(query).
		WithArgs(userID, link).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := s.repo.CreateLink(s.ctx, userID, link)
	require.NoError(s.T(), err)
}

func (s *WishListLinksRepoTestSuite) TestCreateLinkQueryError() {
	userID := uint32(123)
	link := "http://example.com"

	const query = `(?i)INSERT INTO wish_list_links\s*\(user_id,\s*link\)\s*VALUES\s*\(\$1,\s*\$2\)`

	s.dbMock.ExpectExec(query).
		WithArgs(userID, link).
		WillReturnError(errors.New("query error"))

	err := s.repo.CreateLink(s.ctx, userID, link)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "query error")
}

func (s *WishListLinksRepoTestSuite) TestDeleteWishListLinkSuccess() {
	link := "http://example.com"

	const query = `(?i)DELETE\s+FROM\s+wish_list_links\s+WHERE\s+link\s+=\s+\$1`

	s.dbMock.ExpectExec(query).
		WithArgs(link).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err := s.repo.DeleteWishListLink(s.ctx, link)
	require.NoError(s.T(), err)
}

func (s *WishListLinksRepoTestSuite) TestDeleteWishListLinkQueryError() {
	link := "http://example.com"

	const query = `(?i)DELETE\s+FROM\s+wish_list_links\s+WHERE\s+link\s+=\s+\$1`

	s.dbMock.ExpectExec(query).
		WithArgs(link).
		WillReturnError(errors.New("query error"))

	err := s.repo.DeleteWishListLink(s.ctx, link)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "query error")
}

func (s *WishListLinksRepoTestSuite) TestGetUserIDFromLinkSuccess() {
	link := "http://example.com"
	expectedUserID := uint32(123)

	const query = `(?i)SELECT\s+user_id\s+FROM\s+wish_list_links\s+WHERE\s+link\s+=\s+\$1`

	s.dbMock.ExpectQuery(query).
		WithArgs(link).
		WillReturnRows(pgxmock.NewRows([]string{"user_id"}).AddRow(expectedUserID))

	userID, err := s.repo.GetUserIDFromLink(s.ctx, link)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedUserID, userID)
}

func (s *WishListLinksRepoTestSuite) TestGetUserIDFromLinkNotFound() {
	link := "http://example.com"

	const query = `(?i)SELECT\s+user_id\s+FROM\s+wish_list_links\s+WHERE\s+link\s+=\s+\$1`

	s.dbMock.ExpectQuery(query).
		WithArgs(link).
		WillReturnError(errors.New("sql: no rows in result set"))

	userID, err := s.repo.GetUserIDFromLink(s.ctx, link)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "failed to query PostgreSQL: sql: no rows in result set")
	require.Equal(s.T(), uint32(0), userID)
}

func (s *WishListLinksRepoTestSuite) TestGetUserIDFromLinkQueryError() {
	link := "http://example.com"

	const query = `(?i)SELECT\s+user_id\s+FROM\s+wish_list_links\s+WHERE\s+link\s+=\s+\$1`

	s.dbMock.ExpectQuery(query).
		WithArgs(link).
		WillReturnError(errors.New("query error"))

	userID, err := s.repo.GetUserIDFromLink(s.ctx, link)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "failed to query PostgreSQL")
	require.Equal(s.T(), uint32(0), userID)
}

func TestWishListLinksRepo(t *testing.T) {
	suite.Run(t, new(WishListLinksRepoTestSuite))
}
