package address

//import (
//	"context"
//	"errors"
//	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
//	"github.com/jackc/pgx/v5"
//	"github.com/pashagolub/pgxmock/v3"
//	"github.com/stretchr/testify/require"
//	"github.com/stretchr/testify/suite"
//	"io"
//	"log/slog"
//	"testing"
//)
//
//type AddressStoreSuite struct {
//	suite.Suite
//	mock pgxmock.PgxConnIface
//	repo *AddressStore
//}
//
//func (suite *AddressStoreSuite) SetupTest() {
//	mock, err := pgxmock.NewConn()
//	require.NoError(suite.T(), err)
//
//	suite.mock = mock
//	suite.repo = &AddressStore{
//		Db:  mock,
//		Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
//	}
//}
//
//func (suite *AddressStoreSuite) TestGetAddressByProfileID_Success() {
//	ctx := context.Background()
//	flat := "101"
//	profileID := uint32(1)
//
//	rows := pgxmock.NewRows([]string{"id", "city", "street", "house", "flat"}).
//		AddRow(1, "New York", "5th Avenue", "10", "101")
//	suite.mock.ExpectQuery("SELECT id, city, street, house, flat").
//		WithArgs(profileID).
//		WillReturnRows(rows)
//
//	expectedAddr := model.Address{
//		Id:     1,
//		City:   "New York",
//		Street: "5th Avenue",
//		House:  "10",
//		Flat:   &flat,
//	}
//
//	addr, err := suite.repo.GetAddressByProfileID(ctx, profileID)
//	require.NoError(suite.T(), err)
//	require.Equal(suite.T(), expectedAddr, addr)
//}
//
//func (suite *AddressStoreSuite) TestGetAddressByProfileID_NoRows() {
//	ctx := context.Background()
//	profileID := uint32(1)
//
//	suite.mock.ExpectQuery("SELECT id, city, street, house, flat").
//		WithArgs(profileID).
//		WillReturnError(pgx.ErrNoRows)
//
//	addr, err := suite.repo.GetAddressByProfileID(ctx, profileID)
//	require.NoError(suite.T(), err)
//	require.True(suite.T(), addr.IsZero())
//}
//
//func (suite *AddressStoreSuite) TestGetAddressByProfileID_QueryError() {
//	ctx := context.Background()
//	profileID := uint32(1)
//
//	expectedError := errors.New("db error")
//	suite.mock.ExpectQuery("SELECT id, city, street, house, flat").
//		WithArgs(profileID).
//		WillReturnError(expectedError)
//
//	addr, err := suite.repo.GetAddressByProfileID(ctx, profileID)
//	require.Error(suite.T(), err)
//	require.Equal(suite.T(), expectedError, err)
//	require.True(suite.T(), addr.IsZero())
//}
//
//func (suite *AddressStoreSuite) TearDownTest() {
//	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
//}
//
//func TestAddressStoreSuite(t *testing.T) {
//	suite.Run(t, new(AddressStoreSuite))
//}
