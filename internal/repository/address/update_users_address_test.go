package address

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
)

type UpdateUsersAddressTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *AddressStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *UpdateUsersAddressTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &AddressStore{Db: s.dbMock, Log: s.logger}
}

func (s *UpdateUsersAddressTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *UpdateUsersAddressTestSuite) TestUpdateUsersAddressSuccess() {
	addressID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	addressModel := model.Address{
		Text: "123 Main Street, Apartment 4B",
	}

	query := `
		INSERT INTO addresses \(user_id, address\)
		VALUES \(\$1, \$2\)
		ON CONFLICT \(user_id\)
		DO UPDATE SET 
			address = EXCLUDED.address
		RETURNING user_id;
	`

	s.dbMock.ExpectExec(query).WithArgs(addressID, addressModel.Text).WillReturnResult(pgxmock.NewResult("EXECUTE", 1))

	err := s.store.UpdateUsersAddress(s.ctx, addressID, addressModel)
	require.NoError(s.T(), err)
}

func (s *UpdateUsersAddressTestSuite) TestUpdateUsersAddressDatabaseError() {
	addressID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	addressModel := model.Address{
		Text: "456 Elm Street, Unit 12",
	}

	query := `
		INSERT INTO addresses \(user_id, address\)
		VALUES \(\$1, \$2\)
		ON CONFLICT \(user_id\)
		DO UPDATE SET 
			address = EXCLUDED.address
		RETURNING user_id;
	`

	s.dbMock.ExpectExec(query).WithArgs(addressID, addressModel.Text).WillReturnError(errors.New("database error"))

	err := s.store.UpdateUsersAddress(s.ctx, addressID, addressModel)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "database error")
}

func TestUpdateUsersAddressTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateUsersAddressTestSuite))
}
