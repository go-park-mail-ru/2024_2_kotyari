package user

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
	"log/slog"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

type UsersTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *UsersStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *UsersTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &UsersStore{db: s.dbMock, log: s.logger}
}

func (s *UsersTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *UsersTestSuite) TestChangePasswordSuccess() {
	userId := uint32(1)
	newPassword := "new_password"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `UPDATE "users" SET "password" = \$1, "updated_at" = CURRENT_TIMESTAMP WHERE "id" = \$2;`

	s.dbMock.ExpectExec(query).WithArgs(newPassword, userId).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err := s.store.ChangePassword(s.ctx, userId, newPassword)
	require.NoError(s.T(), err)
}

func (s *UsersTestSuite) TestChangePasswordExecError() {
	userId := uint32(1)
	newPassword := "new_password"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `UPDATE "users" SET "password" = \$1, "updated_at" = CURRENT_TIMESTAMP WHERE "id" = \$2;`

	s.dbMock.ExpectExec(query).WithArgs(newPassword, userId).WillReturnError(errors.New("exec error"))

	err := s.store.ChangePassword(s.ctx, userId, newPassword)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "exec error")
}

func (s *UsersTestSuite) TestCreateUserSuccess() {
	userModel := model.User{
		Email:    "test@example.com",
		Username: "testuser",
		Password: "securepassword",
	}
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	insertQuery := `insert into users\(email, username, password\) values \(\$1, \$2, \$3\) returning id, city, username, avatar_url;`

	s.dbMock.ExpectQuery(insertQuery).WithArgs(userModel.Email, userModel.Username, userModel.Password).
		WillReturnRows(pgxmock.NewRows([]string{"id", "city", "username", "avatar_url"}).
			AddRow(uint32(1), "TestCity", "testuser", "avatar.png"))

	createdUser, err := s.store.CreateUser(s.ctx, userModel)
	require.NoError(s.T(), err)
	require.Equal(s.T(), uint32(1), createdUser.ID)
	require.Equal(s.T(), "TestCity", createdUser.City)
	require.Equal(s.T(), "testuser", createdUser.Username)
	require.Equal(s.T(), "avatar.png", createdUser.AvatarUrl)
}

func (s *UsersTestSuite) TestCreateUserAlreadyExists() {
	userModel := model.User{
		Email:    "test@example.com",
		Username: "testuser",
		Password: "securepassword",
	}
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	getUserByEmailQuery := `select id, username, password, city, avatar_url from users where users.email =\$1;`
	s.dbMock.ExpectQuery(getUserByEmailQuery).WithArgs(userModel.Email).
		WillReturnRows(pgxmock.NewRows([]string{"id", "username", "password", "city", "avatar_url"}).
			AddRow(uint32(1), "testuser", "hashedpassword", "TestCity", "avatar.png"))

	_, err := s.store.CreateUser(s.ctx, userModel)
	require.ErrorIs(s.T(), err, errs.UserAlreadyExists)
	err = s.dbMock.ExpectationsWereMet()
	require.NoError(s.T(), err)
}

func (s *UsersTestSuite) TestGetUserByEmailSuccess() {
	userModel := model.User{Email: "test@example.com"}
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select id, username, password, city, avatar_url from users where users.email =\$1;`

	s.dbMock.ExpectQuery(query).WithArgs(userModel.Email).
		WillReturnRows(pgxmock.NewRows([]string{"id", "username", "password", "city", "avatar_url"}).
			AddRow(uint32(1), "testuser", "securepassword", "TestCity", "avatar.png"))

	retrievedUser, err := s.store.GetUserByEmail(s.ctx, userModel)
	require.NoError(s.T(), err)
	require.Equal(s.T(), uint32(1), retrievedUser.ID)
	require.Equal(s.T(), "testuser", retrievedUser.Username)
	require.Equal(s.T(), "TestCity", retrievedUser.City)
	require.Equal(s.T(), "avatar.png", retrievedUser.AvatarUrl)
	err = s.dbMock.ExpectationsWereMet()
	require.NoError(s.T(), err)
}

func (s *UsersTestSuite) TestGetUserByEmailNotFound() {
	userModel := model.User{Email: "notfound@example.com"}
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select id, username, password, city, avatar_url from users where users.email =\$1;`

	s.dbMock.ExpectQuery(query).WithArgs(userModel.Email).WillReturnError(pgx.ErrNoRows)

	_, err := s.store.GetUserByEmail(s.ctx, userModel)
	require.ErrorIs(s.T(), err, errs.UserDoesNotExist)
}

func (s *UsersTestSuite) TestGetUserByUserIDSuccess() {
	userId := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select username, email, city, avatar_url from users where id=\$1;`

	s.dbMock.ExpectQuery(query).WithArgs(userId).
		WillReturnRows(pgxmock.NewRows([]string{"username", "email", "city", "avatar_url"}).
			AddRow("testuser", "test@example.com", "TestCity", "avatar.png"))

	retrievedUser, err := s.store.GetUserByUserID(s.ctx, userId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), "testuser", retrievedUser.Username)
	require.Equal(s.T(), "test@example.com", retrievedUser.Email)
	require.Equal(s.T(), "TestCity", retrievedUser.City)
	require.Equal(s.T(), "avatar.png", retrievedUser.AvatarUrl)
}

func (s *UsersTestSuite) TestGetUserByUserIDNotFound() {
	userId := uint32(999)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select username, email, city, avatar_url from users where id=\$1;`

	s.dbMock.ExpectQuery(query).WithArgs(userId).WillReturnError(pgx.ErrNoRows)

	_, err := s.store.GetUserByUserID(s.ctx, userId)
	require.ErrorIs(s.T(), err, errs.UserDoesNotExist)
}

func TestSearchTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}
