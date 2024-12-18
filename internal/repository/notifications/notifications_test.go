package notifications

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
)

type NotificationsTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *NotificationsStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *NotificationsTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)
	s.logger = logger.InitLogger()
	s.store = &NotificationsStore{
		db:  s.dbMock,
		log: s.logger,
	}
	requestID := "test-request-id"
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)
}

func (s *NotificationsTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *NotificationsTestSuite) TestChangeOrdersStatesSuccess() {
	changeToPaidQuery := `
		update orders
		set new_status = \$1, updated_at = now\(\)
		where status = \$2
		and created_at <= now\(\) - \$3::interval;
	`

	s.dbMock.ExpectExec(changeToPaidQuery).
		WithArgs(Paid, AwaitingPayment, DefaultStateSwitchInterval).
		WillReturnResult(pgxmock.NewResult("UPDATE", 5))

	changeToDeliveredQuery := `
		update orders
		set new_status = \$1
		where status = \$2
		and updated_at <= now\(\) - \$3::interval;
	`

	s.dbMock.ExpectExec(changeToDeliveredQuery).
		WithArgs(Delivered, Paid, DefaultStateSwitchInterval).
		WillReturnResult(pgxmock.NewResult("UPDATE", 3))

	err := s.store.ChangeOrdersStates()
	require.NoError(s.T(), err)
}

func (s *NotificationsTestSuite) TestChangeOrdersStatesQueryError() {
	changeToPaidQuery := `
		update orders
		set new_status = \$1, updated_at = now\(\)
		where status = \$2
		and created_at <= now\(\) - \$3::interval;
	`

	s.dbMock.ExpectExec(changeToPaidQuery).
		WithArgs(Paid, AwaitingPayment, DefaultStateSwitchInterval).
		WillReturnError(errors.New("query error"))

	err := s.store.ChangeOrdersStates()
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "query error")
}

func (s *NotificationsTestSuite) TestGetUserOrdersStatesSuccess() {
	userID := uint32(123)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)
	selectStatusesQuery := `
		select id, new_status
		from orders
		where user_id = \$1
		and status != new_status;
	`

	rows := pgxmock.NewRows([]string{"id", "new_status"}).
		AddRow(uuid.New(), "Paid").
		AddRow(uuid.New(), "Delivered")

	s.dbMock.ExpectQuery(selectStatusesQuery).WithArgs(userID).WillReturnRows(rows)

	updateStatusesQuery := `
		update orders
		set status = new_status
		where user_id = \$1;
	`

	s.dbMock.ExpectExec(updateStatusesQuery).WithArgs(userID).WillReturnResult(pgxmock.NewResult("UPDATE", 2))

	orderStates, err := s.store.GetUserOrdersStates(s.ctx, userID)
	require.NoError(s.T(), err)
	require.Len(s.T(), orderStates, 2)
	require.Equal(s.T(), "Paid", orderStates[0].State)
	require.Equal(s.T(), "Delivered", orderStates[1].State)
}

func (s *NotificationsTestSuite) TestGetUserOrdersStatesNoRows() {
	userID := uint32(123)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	selectStatusesQuery := `
		select id, new_status
		from orders
		where user_id = \$1
		and status != new_status;
	`

	s.dbMock.ExpectQuery(selectStatusesQuery).WithArgs(userID).WillReturnRows(pgxmock.NewRows(nil))

	orderStates, err := s.store.GetUserOrdersStates(s.ctx, userID)
	require.ErrorIs(s.T(), err, errs.NoOrdersUpdates)
	require.Nil(s.T(), orderStates)
}

func (s *NotificationsTestSuite) TestGetUserOrdersStatesQueryError() {
	userID := uint32(123)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	selectStatusesQuery := `
		select id, new_status
		from orders
		where user_id = \$1
		and status != new_status;
	`

	s.dbMock.ExpectQuery(selectStatusesQuery).WithArgs(userID).WillReturnError(errors.New("query error"))

	orderStates, err := s.store.GetUserOrdersStates(s.ctx, userID)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "query error")
	require.Nil(s.T(), orderStates)
}

func TestNotificationsTestSuite(t *testing.T) {
	suite.Run(t, new(NotificationsTestSuite))
}
