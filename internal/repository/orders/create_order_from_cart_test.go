package rorders

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type OrdersRepoSuite struct {
	suite.Suite
	mock  pgxmock.PgxConnIface
	repo  *OrdersRepo
	order *model.OrderFromCart
}

func (suite *OrdersRepoSuite) SetupTest() {
	mock, err := pgxmock.NewConn()
	require.NoError(suite.T(), err)

	suite.mock = mock
	suite.repo = NewOrdersRepo(suite.mock, slog.Default())

	orderID := uuid.New()
	var userID uint32
	faker.FakeData(&userID)

	var id1, id2 uint32
	faker.FakeData(&id1)
	faker.FakeData(&id2)

	var optionID uint32
	faker.FakeData(&optionID)
	suite.order = &model.OrderFromCart{
		OrderID:      orderID,
		UserID:       userID,
		TotalPrice:   1000,
		Address:      "123 Test Street",
		DeliveryDate: time.Now().Add(48 * time.Hour),
		Products: []model.ProductOrder{
			{ID: id1, OptionID: &optionID, Count: 2},
			{ID: id2, Count: 1},
		},
	}
}

func (suite *OrdersRepoSuite) TestCreateOrderFromCart() {
	ctx := context.Background()

	suite.mock.ExpectBegin()

	suite.mock.ExpectQuery(`INSERT INTO orders\s*\(id, user_id, total_price, address, created_at, updated_at\)\s*VALUES\s*\(\$1, \$2, \$3, \$4, NOW\(\), NOW\(\)\)\s*RETURNING created_at;`).
		WithArgs(suite.order.OrderID, suite.order.UserID, suite.order.TotalPrice, suite.order.Address).
		WillReturnRows(pgxmock.NewRows([]string{"created_at"}).AddRow(time.Now()))

	for _, p := range suite.order.Products {
		suite.mock.ExpectExec(`INSERT INTO product_orders\s*\(id, order_id, product_id, option_id, count, delivery_date\)\s*VALUES\s*\(\$1, \$2, \$3, \$4, \$5, \$6\);`).
			WithArgs(pgxmock.AnyArg(), suite.order.OrderID, p.ID, p.OptionID, p.Count, suite.order.DeliveryDate).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))
	}

	suite.mock.ExpectExec(`UPDATE carts`).
		WithArgs(suite.order.UserID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	suite.mock.ExpectCommit()

	order, err := suite.repo.CreateOrderFromCart(ctx, suite.order)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), order)
	require.Equal(suite.T(), suite.order.OrderID, order.ID)
	require.Equal(suite.T(), suite.order.Address, order.Address)
	require.Equal(suite.T(), suite.order.TotalPrice, order.TotalPrice)
	require.Equal(suite.T(), suite.order.Products, order.Products)
}

func (suite *OrdersRepoSuite) TearDownTest() {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func TestOrdersRepoSuite(t *testing.T) {
	suite.Run(t, new(OrdersRepoSuite))
}
