package profile

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

type UpdateProfileTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *ProfilesStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *UpdateProfileTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &ProfilesStore{db: s.dbMock, log: s.logger}
}

func (s *UpdateProfileTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *UpdateProfileTestSuite) TestUpdateProfileSuccess() {
	profileID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		UPDATE users 
		SET email = \$1, 
		    username = \$2, 
		    gender = \$3 
		WHERE id = \$4;
	`

	profileToUpdate := model.Profile{
		Email:    "updated@example.com",
		Username: "updateduser",
		Gender:   "female",
	}

	s.dbMock.ExpectExec(query).
		WithArgs(profileToUpdate.Email, profileToUpdate.Username, profileToUpdate.Gender, profileID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err := s.store.UpdateProfile(s.ctx, profileID, profileToUpdate)
	require.NoError(s.T(), err)
}

func (s *UpdateProfileTestSuite) TestUpdateProfileNoRowsUpdated() {
	profileID := uint32(999)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		UPDATE users 
		SET email = \$1, 
		    username = \$2, 
		    gender = \$3 
		WHERE id = \$4;
	`

	profileToUpdate := model.Profile{
		Email:    "nonexistent@example.com",
		Username: "nonexistentuser",
		Gender:   "male",
	}

	s.dbMock.ExpectExec(query).
		WithArgs(profileToUpdate.Email, profileToUpdate.Username, profileToUpdate.Gender, profileID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 0))

	err := s.store.UpdateProfile(s.ctx, profileID, profileToUpdate)
	require.NoError(s.T(), err)
}

func (s *UpdateProfileTestSuite) TestUpdateProfileExecError() {
	profileID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		UPDATE users 
		SET email = \$1, 
		    username = \$2, 
		    gender = \$3 
		WHERE id = \$4;
	`

	profileToUpdate := model.Profile{
		Email:    "updated@example.com",
		Username: "updateduser",
		Gender:   "female",
	}

	s.dbMock.ExpectExec(query).
		WithArgs(profileToUpdate.Email, profileToUpdate.Username, profileToUpdate.Gender, profileID).
		WillReturnError(errors.New("database error"))

	err := s.store.UpdateProfile(s.ctx, profileID, profileToUpdate)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "database error")
}

func TestUpdateProfileTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateProfileTestSuite))
}
