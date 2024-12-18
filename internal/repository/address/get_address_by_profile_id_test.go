package address

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
)

type GetAddressByProfileIDTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *AddressStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *GetAddressByProfileIDTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &AddressStore{Db: s.dbMock, Log: s.logger}
}

func (s *GetAddressByProfileIDTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *GetAddressByProfileIDTestSuite) TestGetAddressByProfileIDSuccess() {
	profileID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		SELECT address
		FROM addresses 
		WHERE addresses.user_id = \$1;
	`

	expectedAddress := model.Address(model.Address{Text: "123 Main Street, Apartment 4B"})
	expectedAddressRow := "123 Main Street, Apartment 4B"
	rows := pgxmock.NewRows([]string{"address"}).AddRow(expectedAddressRow)

	s.dbMock.ExpectQuery(query).WithArgs(profileID).WillReturnRows(rows)

	address, err := s.store.GetAddressByProfileID(s.ctx, profileID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedAddress, address)
}

func (s *GetAddressByProfileIDTestSuite) TestGetAddressByProfileIDNotFound() {
	profileID := uint32(999)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		SELECT address
		FROM addresses 
		WHERE addresses.user_id = \$1;
	`

	s.dbMock.ExpectQuery(query).WithArgs(profileID).WillReturnError(pgx.ErrNoRows)

	address, err := s.store.GetAddressByProfileID(s.ctx, profileID)
	require.NoError(s.T(), err)
	require.Empty(s.T(), address)
}

func (s *GetAddressByProfileIDTestSuite) TestGetAddressByProfileIDQueryError() {
	profileID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		SELECT address
		FROM addresses 
		WHERE addresses.user_id = \$1;
	`

	s.dbMock.ExpectQuery(query).WithArgs(profileID).WillReturnError(errors.New("database error"))

	address, err := s.store.GetAddressByProfileID(s.ctx, profileID)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "database error")
	require.Empty(s.T(), address)
}

func TestGetAddressByProfileIDTestSuite(t *testing.T) {
	suite.Run(t, new(GetAddressByProfileIDTestSuite))
}
