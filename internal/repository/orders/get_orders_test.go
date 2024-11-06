package rorders

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type OrdersRepoGetOrdersSuite struct {
	suite.Suite
	mock pgxmock.PgxConnIface
	repo *OrdersRepo
}

func (suite *OrdersRepoGetOrdersSuite) SetupTest() {
	mock, err := pgxmock.NewConn()
	require.NoError(suite.T(), err)

	suite.mock = mock
	suite.repo = NewOrdersRepo(suite.mock, slog.Default())
}

func (suite *OrdersRepoGetOrdersSuite) TestGetOrders_Success() {
	ctx := context.Background()
	userID := uint32(12345)

	orderID := uuid.New()
	orderDate := time.Now().Add(-24 * time.Hour)
	deliveryDate := time.Now().Add(24 * time.Hour)

	rows := pgxmock.NewRows([]string{"id", "order_date", "delivery_date", "product_id", "image_url", "name", "total_price", "status"}).
		AddRow(orderID, orderDate, deliveryDate, uint32(1), "image1.jpg", "Product 1", uint16(500), "Pending").
		AddRow(orderID, orderDate, deliveryDate, uint32(2), "image2.jpg", "Product 2", uint16(500), "Pending")

	suite.mock.ExpectQuery(`SELECT o.id::uuid, o.created_at AS order_date, po.delivery_date, p.id::bigint AS product_id, p.image_url, p.title AS name, o.total_price, o.status FROM orders o JOIN product_orders po ON o.id = po.order_id JOIN products p ON po.product_id = p.id WHERE o.user_id = \$1 ORDER BY po.delivery_date, o.created_at;`).
		WithArgs(userID).
		WillReturnRows(rows)

	orders, err := suite.repo.GetOrders(ctx, userID)
	require.NoError(suite.T(), err)
	require.Len(suite.T(), orders, 1)
	require.Equal(suite.T(), orderID, orders[0].ID)
	require.Equal(suite.T(), deliveryDate, orders[0].DeliveryDate)
	require.Equal(suite.T(), uint16(500), orders[0].TotalPrice)
	require.Equal(suite.T(), "Pending", orders[0].Status)
	require.Len(suite.T(), orders[0].Products, 2)
	require.Equal(suite.T(), uint32(1), orders[0].Products[0].ProductID)
	require.Equal(suite.T(), "Product 1", orders[0].Products[0].Name)
	require.Equal(suite.T(), uint32(2), orders[0].Products[1].ProductID)
	require.Equal(suite.T(), "Product 2", orders[0].Products[1].Name)
}

func (suite *OrdersRepoGetOrdersSuite) TestGetOrders_QueryError() {
	ctx := context.Background()
	userID := uint32(12345)

	expectedError := errors.New("database error")
	suite.mock.ExpectQuery(`SELECT o.id::uuid, o.created_at AS order_date, po.delivery_date, p.id::bigint AS product_id, p.image_url, p.title AS name, o.total_price, o.status FROM orders o JOIN product_orders po ON o.id = po.order_id JOIN products p ON po.product_id = p.id WHERE o.user_id = \$1 ORDER BY po.delivery_date, o.created_at;`).
		WithArgs(userID).
		WillReturnError(expectedError)

	orders, err := suite.repo.GetOrders(ctx, userID)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), orders)
	require.Equal(suite.T(), expectedError, err)
}

func (suite *OrdersRepoGetOrdersSuite) TearDownTest() {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func TestOrdersRepoGetOrdersSuite(t *testing.T) {
	suite.Run(t, new(OrdersRepoGetOrdersSuite))
}
