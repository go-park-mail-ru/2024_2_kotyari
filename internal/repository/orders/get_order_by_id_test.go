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

type OrdersRepoGetOrderByIdSuite struct {
	suite.Suite
	mock pgxmock.PgxConnIface
	repo *OrdersRepo
}

func (suite *OrdersRepoGetOrderByIdSuite) SetupTest() {
	mock, err := pgxmock.NewConn()
	require.NoError(suite.T(), err)

	suite.mock = mock
	suite.repo = NewOrdersRepo(suite.mock, slog.Default())
}

func (suite *OrdersRepoGetOrderByIdSuite) TestGetOrderById_Success() {
	ctx := context.Background()
	userID := uint32(12345)
	orderID := uuid.New()
	deliveryDate := time.Now()

	rows := pgxmock.NewRows([]string{"id", "address", "status", "created_at", "username", "delivery_date", "product_id", "price", "count", "image_url", "weight", "title"}).
		AddRow(orderID, "123 Main St", "Delivered", deliveryDate.Add(-48*time.Hour), "john_doe", deliveryDate, uint32(1), uint16(1000), uint16(2), "image1.jpg", uint16(2), "Product A").
		AddRow(orderID, "123 Main St", "Delivered", deliveryDate.Add(-48*time.Hour), "john_doe", deliveryDate, uint32(2), uint16(500), uint16(1), "image2.jpg", uint16(3), "Product B")

	suite.mock.ExpectQuery(`SELECT o.id, o.address, o.status, o.created_at, u.username, op.delivery_date, p.id, p.price, op.count, p.image_url, p.weight, p.title FROM orders o JOIN users u ON o.user_id = u.id JOIN product_orders op ON o.id = op.order_id JOIN products p ON op.product_id = p.id WHERE o.id = \$1::uuid AND o.user_id = \$2 AND op.delivery_date BETWEEN \$3 AND \$4;`).
		WithArgs(orderID, userID, deliveryDate.Truncate(time.Millisecond), deliveryDate.Truncate(time.Millisecond).Add(time.Millisecond)).
		WillReturnRows(rows)

	order, err := suite.repo.GetOrderById(ctx, orderID, userID)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), order)
	require.Equal(suite.T(), orderID, order.ID)
	require.Equal(suite.T(), "123 Main St", order.Address)
	require.Equal(suite.T(), "Delivered", order.Status)
	require.Len(suite.T(), order.Products, 2)

	require.Equal(suite.T(), uint32(1), order.Products[0].ProductID)
	require.Equal(suite.T(), uint16(1000), order.Products[0].Cost)
	require.Equal(suite.T(), "Product A", order.Products[0].Name)

	require.Equal(suite.T(), uint32(2), order.Products[1].ProductID)
	require.Equal(suite.T(), uint16(500), order.Products[1].Cost)
	require.Equal(suite.T(), "Product B", order.Products[1].Name)
}

func (suite *OrdersRepoGetOrderByIdSuite) TestGetOrderById_QueryError() {
	ctx := context.Background()
	userID := uint32(12345)
	orderID := uuid.New()
	deliveryDate := time.Now()

	expectedError := errors.New("database error")
	suite.mock.ExpectQuery(`SELECT o.id, o.address, o.status, o.created_at, u.username, op.delivery_date, p.id, p.price, op.count, p.image_url, p.weight, p.title FROM orders o JOIN users u ON o.user_id = u.id JOIN product_orders op ON o.id = op.order_id JOIN products p ON op.product_id = p.id WHERE o.id = \$1::uuid AND o.user_id = \$2 AND op.delivery_date BETWEEN \$3 AND \$4;`).
		WithArgs(orderID, userID, deliveryDate.Truncate(time.Millisecond), deliveryDate.Truncate(time.Millisecond).Add(time.Millisecond)).
		WillReturnError(expectedError)

	order, err := suite.repo.GetOrderById(ctx, orderID, userID, deliveryDate)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), order)
	require.Equal(suite.T(), expectedError, err)
}

func (suite *OrdersRepoGetOrderByIdSuite) TearDownTest() {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func TestOrdersRepoGetOrderByIdSuite(t *testing.T) {
	suite.Run(t, new(OrdersRepoGetOrderByIdSuite))
}
