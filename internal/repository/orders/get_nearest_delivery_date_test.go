package rorders

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type OrdersRepoGetNearestDeliveryDateSuite struct {
	suite.Suite
	mock pgxmock.PgxConnIface
	repo *OrdersRepo
}

func (suite *OrdersRepoGetNearestDeliveryDateSuite) SetupTest() {
	mock, err := pgxmock.NewConn()
	require.NoError(suite.T(), err)

	suite.mock = mock
	suite.repo = NewOrdersRepo(suite.mock, slog.Default())
}

func (suite *OrdersRepoGetNearestDeliveryDateSuite) TestGetNearestDeliveryDate_Success() {
	ctx := context.Background()
	var userID uint32 = 12345

	expectedDate := time.Now().Add(24 * time.Hour)
	suite.mock.ExpectQuery(`SELECT MIN\(po.delivery_date\) FROM orders o JOIN product_orders po ON o.id = po.order_id WHERE o.user_id = \$1 AND po.delivery_date > NOW\(\);`).
		WithArgs(userID).
		WillReturnRows(pgxmock.NewRows([]string{"delivery_date"}).AddRow(expectedDate))

	deliveryDate, err := suite.repo.GetNearestDeliveryDate(ctx, userID)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), expectedDate, deliveryDate)
}

func (suite *OrdersRepoGetNearestDeliveryDateSuite) TestGetNearestDeliveryDate_NoRows() {
	ctx := context.Background()
	var userID uint32 = 12345

	suite.mock.ExpectQuery(`SELECT MIN\(po.delivery_date\) FROM orders o JOIN product_orders po ON o.id = po.order_id WHERE o.user_id = \$1 AND po.delivery_date > NOW\(\);`).
		WithArgs(userID).
		WillReturnError(pgx.ErrNoRows)

	deliveryDate, err := suite.repo.GetNearestDeliveryDate(ctx, userID)
	require.NoError(suite.T(), err)
	require.True(suite.T(), deliveryDate.IsZero())
}

func (suite *OrdersRepoGetNearestDeliveryDateSuite) TestGetNearestDeliveryDate_QueryError() {
	ctx := context.Background()
	var userID uint32 = 12345

	expectedError := errors.New("database error")
	suite.mock.ExpectQuery(`SELECT MIN\(po.delivery_date\) FROM orders o JOIN product_orders po ON o.id = po.order_id WHERE o.user_id = \$1 AND po.delivery_date > NOW\(\);`).
		WithArgs(userID).
		WillReturnError(expectedError)

	deliveryDate, err := suite.repo.GetNearestDeliveryDate(ctx, userID)
	require.Error(suite.T(), err)
	require.Equal(suite.T(), expectedError, err)
	require.True(suite.T(), deliveryDate.IsZero())
}

func (suite *OrdersRepoGetNearestDeliveryDateSuite) TearDownTest() {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func TestOrdersRepoGetNearestDeliveryDateSuite(t *testing.T) {
	suite.Run(t, new(OrdersRepoGetNearestDeliveryDateSuite))
}
