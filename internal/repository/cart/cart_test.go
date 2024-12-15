package cart

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"testing"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
)

type CartsStoreTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *CartsStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *CartsStoreTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &CartsStore{db: s.dbMock, log: s.logger}
}

func (s *CartsStoreTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *CartsStoreTestSuite) TestAddProductSuccess() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	insertQuery := `
		insert into carts\(user_id, product_id, count, is_selected, is_deleted\)
		values \(\$1, \$2, 1, true, false\)
	`
	updateProductCountQuery := `^update products set count=count-\$2 where id=\$1;$`

	s.dbMock.ExpectBeginTx(pgx.TxOptions{AccessMode: pgx.ReadWrite})
	s.dbMock.ExpectExec(insertQuery).WithArgs(userID, productID).WillReturnResult(pgxmock.NewResult("INSERT", 1))
	s.dbMock.ExpectExec(updateProductCountQuery).WithArgs(productID, int32(1)).WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	s.dbMock.ExpectCommit()
	err := s.store.AddProduct(s.ctx, productID, userID)
	require.NoError(s.T(), err)
}

func (s *CartsStoreTestSuite) TestAddProductQueryError() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	insertQuery := `
		insert into carts\(user_id, product_id, count, is_selected, is_deleted\)
		values \(\$1, \$2, 1, true, false\)
	`
	s.dbMock.ExpectBeginTx(pgx.TxOptions{AccessMode: pgx.ReadWrite})
	s.dbMock.ExpectExec(insertQuery).WithArgs(userID, productID).WillReturnError(errors.New("database error"))
	s.dbMock.ExpectRollback()

	err := s.store.AddProduct(s.ctx, productID, userID)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "database error")
}

func (s *CartsStoreTestSuite) TestChangeAllCartProductsStateSuccess() {
	userID := uint32(2)
	isSelected := true
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	updateQuery := `
		update carts
		set is_selected = \$2
		where user_id = \$1 and is_deleted = false;
	`
	s.dbMock.ExpectExec(updateQuery).WithArgs(userID, isSelected).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err := s.store.ChangeAllCartProductsState(s.ctx, userID, isSelected)
	require.NoError(s.T(), err)
}

func (s *CartsStoreTestSuite) TestChangeAllCartProductsStateNoRows() {
	userID := uint32(2)
	isSelected := true
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	updateQuery := `
		update carts
		set is_selected = \$2
		where user_id = \$1 and is_deleted = false;
	`
	s.dbMock.ExpectExec(updateQuery).WithArgs(userID, isSelected).WillReturnResult(pgxmock.NewResult("UPDATE", 0))

	err := s.store.ChangeAllCartProductsState(s.ctx, userID, isSelected)
	require.ErrorIs(s.T(), err, errs.EmptyCart)
}

func (s *CartsStoreTestSuite) TestChangeCartProductCountSuccess() {
	productID := uint32(1)
	count := int32(3)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	updateProductCountQuery := `^update products set count=count-\$2 where id=\$1;$`
	updateCartCountQuery := `^update carts set count=count\+\$2 where product_id=\$1 and user_id=\$3;$`

	s.dbMock.ExpectBeginTx(pgx.TxOptions{AccessMode: pgx.ReadWrite})
	s.dbMock.ExpectExec(updateProductCountQuery).WithArgs(productID, count).WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	s.dbMock.ExpectExec(updateCartCountQuery).WithArgs(productID, count, userID).WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	s.dbMock.ExpectCommit()

	err := s.store.ChangeCartProductCount(s.ctx, productID, count, userID)
	require.NoError(s.T(), err)
}

func (s *CartsStoreTestSuite) TestChangeCartProductCountQueryError() {
	productID := uint32(1)
	count := int32(3)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	updateProductCountQuery := `^update products set count=count-\$2 where id=\$1;$`
	s.dbMock.ExpectBeginTx(pgx.TxOptions{AccessMode: pgx.ReadWrite})
	s.dbMock.ExpectExec(updateProductCountQuery).WithArgs(productID, count).WillReturnError(errors.New("database error"))
	s.dbMock.ExpectRollback()

	err := s.store.ChangeCartProductCount(s.ctx, productID, count, userID)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "database error")
}

func TestCartsStoreTestSuite(t *testing.T) {
	suite.Run(t, new(CartsStoreTestSuite))
}
