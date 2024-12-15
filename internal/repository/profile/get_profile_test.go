package profile

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
)

type ProfileTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *ProfilesStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *ProfileTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &ProfilesStore{db: s.dbMock, log: s.logger}
}

func (s *ProfileTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *ProfileTestSuite) TestGetProfileSuccess() {
	userID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `SELECT id, email, username, age, gender, avatar_url FROM users WHERE users\.id = \$1;`

	expectedProfile := model.Profile{
		ID:        userID,
		Email:     "user@example.com",
		Username:  "testuser",
		Age:       25,
		Gender:    "male",
		AvatarURL: "https://example.com/avatar.png",
	}

	rows := pgxmock.NewRows([]string{"id", "email", "username", "age", "gender", "avatar_url"}).
		AddRow(expectedProfile.ID, expectedProfile.Email, expectedProfile.Username, expectedProfile.Age, expectedProfile.Gender, expectedProfile.AvatarURL)

	s.dbMock.ExpectQuery(query).WithArgs(userID).WillReturnRows(rows)

	profile, err := s.store.GetProfile(s.ctx, userID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedProfile, profile)
}

func (s *ProfileTestSuite) TestGetProfileUserNotFound() {
	userID := uint32(999)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `SELECT id, email, username, age, gender, avatar_url FROM users WHERE users\.id = \$1;`

	s.dbMock.ExpectQuery(query).WithArgs(userID).WillReturnError(pgx.ErrNoRows)

	profile, err := s.store.GetProfile(s.ctx, userID)
	require.Error(s.T(), err)
	require.True(s.T(), errors.Is(err, errs.UserDoesNotExist))
	require.Equal(s.T(), model.Profile{}, profile)
}

func (s *ProfileTestSuite) TestGetProfileQueryError() {
	userID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `SELECT id, email, username, age, gender, avatar_url FROM users WHERE users\.id = \$1;`

	s.dbMock.ExpectQuery(query).WithArgs(userID).WillReturnError(errors.New("query error"))

	profile, err := s.store.GetProfile(s.ctx, userID)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "query error")
	require.Equal(s.T(), model.Profile{}, profile)
}

func TestProfileTestSuite(t *testing.T) {
	suite.Run(t, new(ProfileTestSuite))
}
