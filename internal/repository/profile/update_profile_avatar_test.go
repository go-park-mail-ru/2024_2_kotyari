package profile

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
)

type UpdateProfileAvatarTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *ProfilesStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *UpdateProfileAvatarTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &ProfilesStore{db: s.dbMock, log: s.logger}
}

func (s *UpdateProfileAvatarTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *UpdateProfileAvatarTestSuite) TestUpdateProfileAvatarSuccess() {
	profileID := uint32(1)
	filePath := "/avatars/user1.png"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		UPDATE users 
		SET avatar_url = \$1
		WHERE id = \$2;
	`

	s.dbMock.ExpectExec(query).
		WithArgs(filePath, profileID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err := s.store.UpdateProfileAvatar(s.ctx, profileID, filePath)
	require.NoError(s.T(), err)
}

func (s *UpdateProfileAvatarTestSuite) TestUpdateProfileAvatarNoRowsUpdated() {
	profileID := uint32(999)
	filePath := "/avatars/nonexistent.png"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		UPDATE users 
		SET avatar_url = \$1
		WHERE id = \$2;
	`

	s.dbMock.ExpectExec(query).
		WithArgs(filePath, profileID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 0))

	err := s.store.UpdateProfileAvatar(s.ctx, profileID, filePath)
	require.NoError(s.T(), err)
}

func (s *UpdateProfileAvatarTestSuite) TestUpdateProfileAvatarExecError() {
	profileID := uint32(1)
	filePath := "/avatars/error.png"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		UPDATE users 
		SET avatar_url = \$1
		WHERE id = \$2;
	`

	s.dbMock.ExpectExec(query).
		WithArgs(filePath, profileID).
		WillReturnError(errors.New("database error"))

	err := s.store.UpdateProfileAvatar(s.ctx, profileID, filePath)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "database error")
}

func TestUpdateProfileAvatarTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateProfileAvatarTestSuite))
}
