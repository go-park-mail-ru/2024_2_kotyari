package cart

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/jackc/pgx/v5"
	"regexp"
	"testing"
	"time"

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

func (s *CartsStoreTestSuite) SetupTest() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &CartsStore{db: s.dbMock, log: s.logger}
	s.ctx = context.Background()
}

func (s *CartsStoreTestSuite) TearDownTest() {
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

func (s *CartsStoreTestSuite) TestChangeCartProductDeletedStateSuccess() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	updateQuery := `
		update carts 
		set is_deleted = false, is_selected = true
		where product_id = \$1 and user_id = \$2;
	`

	incrementProductCountQuery := `^update products set count=count-\$2 where id=\$1;$`
	incrementCartCountQuery := `^update carts set count=count\+\$2 where product_id=\$1 and user_id=\$3;$`

	s.dbMock.ExpectExec(updateQuery).WithArgs(productID, userID).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	s.dbMock.ExpectBeginTx(pgx.TxOptions{AccessMode: pgx.ReadWrite})
	s.dbMock.ExpectExec(incrementProductCountQuery).WithArgs(productID, int32(1)).WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	s.dbMock.ExpectExec(incrementCartCountQuery).WithArgs(productID, int32(1), userID).WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	s.dbMock.ExpectCommit()

	err := s.store.ChangeCartProductDeletedState(s.ctx, productID, userID)
	require.NoError(s.T(), err)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *CartsStoreTestSuite) TestChangeCartProductDeletedStateNoRowsAffected() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	updateQuery := `
		update carts 
		set is_deleted = false, is_selected = true
		where product_id = \$1 and user_id = \$2;
	`

	s.dbMock.ExpectExec(updateQuery).WithArgs(productID, userID).WillReturnResult(pgxmock.NewResult("UPDATE", 0))

	err := s.store.ChangeCartProductDeletedState(s.ctx, productID, userID)
	require.ErrorIs(s.T(), err, errs.ProductNotInCart)
}

func (s *CartsStoreTestSuite) TestChangeCartProductSelectedStateSuccess() {
	productID := uint32(1)
	userID := uint32(2)
	isSelected := true
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	updateQuery := `
		update carts 
		set is_selected = \$3
		where product_id = \$1 and user_id = \$2;
	`

	s.dbMock.ExpectExec(updateQuery).WithArgs(productID, userID, isSelected).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err := s.store.ChangeCartProductSelectedState(s.ctx, productID, userID, isSelected)
	require.NoError(s.T(), err)
}

func (s *CartsStoreTestSuite) TestChangeCartProductSelectedStateNoRowsAffected() {
	productID := uint32(1)
	userID := uint32(2)
	isSelected := true
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	updateQuery := `
		update carts 
		set is_selected = \$3
		where product_id = \$1 and user_id = \$2;
	`

	s.dbMock.ExpectExec(updateQuery).WithArgs(productID, userID, isSelected).WillReturnResult(pgxmock.NewResult("UPDATE", 0))

	err := s.store.ChangeCartProductSelectedState(s.ctx, productID, userID, isSelected)
	require.ErrorIs(s.T(), err, errs.ProductNotInCart)
}

func (s *CartsStoreTestSuite) TestGetCartSuccess() {
	userID := uint32(2)
	deliveryDate := time.Now()
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	selectQuery := `
	select c.id, p.id, title, price, p.weight, description, image_url, original_price, discount, c.count, c.is_selected 
	from products p
	join carts c on p.id = c.product_id 
	where user_id=\$1 and c.is_deleted = false and c.count>0;
`

	rows := pgxmock.NewRows([]string{
		"id", "product_id", "title", "price", "weight", "description", "image_url", "original_price", "discount", "count", "is_selected",
	}).AddRow(
		uint32(1), uint32(1), "Test Product", uint32(100), float32(2.0), "Test Description", "http://example.com/image.jpg", uint32(120), uint32(20), uint32(2), true,
	)

	s.dbMock.ExpectQuery(selectQuery).WithArgs(userID).WillReturnRows(rows)

	cart, err := s.store.GetCart(s.ctx, userID, deliveryDate)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), cart)
	require.Equal(s.T(), userID, cart.UserID)
	require.Len(s.T(), cart.Products, 1)
}

func (s *CartsStoreTestSuite) TestGetCartNoRows() {
	userID := uint32(2)
	deliveryDate := time.Now()
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	selectQuery := `
		select c.id, p.id, title, price, p.weight, description, image_url, original_price, discount, c.count, c.is_selected  
		from products p
		join carts c on p.id = c.product_id 
		where user_id=\$1 and c.is_deleted = false and c.count>0;
	`

	s.dbMock.ExpectQuery(selectQuery).WithArgs(userID).WillReturnError(pgx.ErrNoRows)

	_, err := s.store.GetCart(s.ctx, userID, deliveryDate)
	require.ErrorIs(s.T(), err, errs.CartDoesNotExist)
}

func (s *CartsStoreTestSuite) TestGetCartProductSuccess() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		select c.count, c.is_selected, c.is_deleted from carts c
		join products p on p.id = c.product_id
		where p.id = \$1 and c.user_id = \$2;$`

	s.dbMock.ExpectQuery(query).
		WithArgs(productID, userID).
		WillReturnRows(pgxmock.NewRows([]string{"count", "is_selected", "is_deleted"}).
			AddRow(3, true, false))

	cartProduct, err := s.store.GetCartProduct(s.ctx, productID, userID)

	require.NoError(s.T(), err)
	require.Equal(s.T(), model.CartProduct{
		IsSelected: false,
		IsDeleted:  false,
	}, cartProduct)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *CartsStoreTestSuite) TestGetCartProductNotFound() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		select c.count, c.is_selected, c.is_deleted from carts c
		join products p on p.id = c.product_id
		where p.id = \$1 and c.user_id = \$2;$`

	s.dbMock.ExpectQuery(query).
		WithArgs(productID, userID).
		WillReturnError(pgx.ErrNoRows)

	_, err := s.store.GetCartProduct(s.ctx, productID, userID)

	require.ErrorIs(s.T(), err, errs.ProductNotInCart)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *CartsStoreTestSuite) TestGetSelectedCartItemsSuccess() {
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		SELECT p.id, p.title, p.image_url, p.price, c.count, p.weight
		FROM carts AS c
		INNER JOIN products AS p ON c.product_id = p.id
		WHERE c.user_id = \$1 AND c.is_selected = true AND c.is_deleted = false;$`

	rows := pgxmock.NewRows([]string{"id", "title", "image_url", "price", "count", "weight"}).
		AddRow(uint32(1), "Product 1", "image1.jpg", uint32(100), uint32(2), float32(1.5)).
		AddRow(uint32(2), "Product 2", "image2.jpg", uint32(200), uint32(1), float32(2.0))

	s.dbMock.ExpectQuery(query).
		WithArgs(userID).
		WillReturnRows(rows)

	items, err := s.store.GetSelectedCartItems(s.ctx, userID)

	require.NoError(s.T(), err)
	require.Len(s.T(), items, 2)
	require.Equal(s.T(), items[0], model.ProductOrder{
		ID:       1,
		Name:     "Product 1",
		ImageUrl: "image1.jpg",
		Cost:     100,
		Count:    2,
		Weight:   1.5,
	})
	require.Equal(s.T(), items[1], model.ProductOrder{
		ID:       2,
		Name:     "Product 2",
		ImageUrl: "image2.jpg",
		Cost:     200,
		Count:    1,
		Weight:   2.0,
	})

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

//
//func (s *CartsStoreTestSuite) TestGetSelectedFromCartSuccess() {
//	userID := uint32(2)
//	requestID := uuid.New()
//	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)
//
//	fixedDeliveryDate := time.Date(2024, time.December, 19, 16, 10, 13, 974805439, time.UTC)
//
//	pgxDeliveryDate := pgtype.Timestamptz{
//		Time:  fixedDeliveryDate,
//		Valid: true,
//	}
//
//	query := `
//		SELECT p.id, p.title, p.price, p.image_url, p.weight, c.count, c.delivery_date, u.username, u.preferred_payment_method,
//		       a.address
//		FROM users u
//			LEFT JOIN carts c ON u.id = c.user_id AND c.is_deleted = false
//			LEFT JOIN products p ON p.id = c.product_id
//			LEFT JOIN addresses a ON u.id = a.user_id
//		WHERE u.id=\$1;$`
//
//	rows := pgxmock.NewRows([]string{"id", "title", "price", "image_url", "weight", "count", "delivery_date", "username", "preferred_payment_method", "address"}).
//		AddRow(
//			pgtype.Int4{Int32: 1, Valid: true},
//			pgtype.Text{String: "Product 1", Valid: true},
//			pgtype.Uint32{Uint32: 100, Valid: true},
//			pgtype.Text{String: "image1.jpg", Valid: true},
//			pgtype.Float4{Float32: 1.5, Valid: true},
//			pgtype.Uint32{Uint32: 2, Valid: true},
//			pgxDeliveryDate,
//			"JohnDoe",
//			"CreditCard",
//			"123 Street",
//		)
//
//	s.dbMock.ExpectQuery(query).
//		WithArgs(userID).
//		WillReturnRows(rows)
//
//	cart, err := s.store.GetSelectedFromCart(s.ctx, userID)
//	require.NoError(s.T(), err)
//	require.Equal(s.T(), "JohnDoe", cart.UserName)
//	require.Equal(s.T(), "CreditCard", cart.PreferredPaymentMethod)
//	require.Equal(s.T(), "123 Street", cart.Address.Text)
//	require.Len(s.T(), cart.Items, 1)
//	require.Equal(s.T(), model.CartProductForOrder{
//		URL:          "/catalog/product/1",
//		Title:        "Product 1",
//		Price:        100,
//		Image:        "image1.jpg",
//		Weight:       1.5,
//		Quantity:     2,
//		DeliveryDate: fixedDeliveryDate,
//	}, cart.Items[0])
//
//	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
//}

func (s *CartsStoreTestSuite) TestProductInCartSuccess() {
	userID := uint32(1)
	productID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := regexp.QuoteMeta(`
        SELECT EXISTS (
            SELECT 1 
            FROM carts 
            WHERE user_id = $1 AND product_id = $2 AND is_deleted = false
        )
    `)

	expectedExists := true

	row := pgxmock.NewRows([]string{"exists"}).AddRow(expectedExists)
	s.dbMock.ExpectQuery(query).WithArgs(userID, productID).WillReturnRows(row)

	exists, err := s.store.ProductInCart(s.ctx, userID, productID)
	require.NoError(s.T(), err)
	require.True(s.T(), exists)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *CartsStoreTestSuite) TestProductInCartError() {
	userID := uint32(1)
	productID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
        SELECT EXISTS (
            SELECT 1 
            FROM carts 
            WHERE user_id = \$1 AND product_id = \$2 AND is_deleted = false
        )
    `

	s.dbMock.ExpectQuery(query).WithArgs(userID, productID).WillReturnError(fmt.Errorf("some error"))

	exists, err := s.store.ProductInCart(s.ctx, userID, productID)
	require.Error(s.T(), err)
	require.False(s.T(), exists)
}

func (s *CartsStoreTestSuite) TestRemoveCartProductSuccess() {
	productID := uint32(1)
	userID := uint32(2)
	count := int32(3)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	deleteQuery := `
        update carts
        set count = 0, is_deleted = true, is_selected = false
        where product_id = \$1 and user_id = \$2;
    `
	updateProductCountQuery := `^update products set count=count-\$2 where id=\$1;$`

	s.dbMock.ExpectBeginTx(pgx.TxOptions{AccessMode: pgx.ReadWrite})
	s.dbMock.ExpectExec(deleteQuery).WithArgs(productID, userID).WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	s.dbMock.ExpectExec(updateProductCountQuery).WithArgs(productID, count).WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	s.dbMock.ExpectCommit()

	err := s.store.RemoveCartProduct(s.ctx, productID, count, userID)
	require.NoError(s.T(), err)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *CartsStoreTestSuite) TestRemoveCartProductTransactionError() {
	productID := uint32(1)
	userID := uint32(2)
	count := int32(3)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	s.dbMock.ExpectBegin().WillReturnError(fmt.Errorf("failed to start transaction"))

	err := s.store.RemoveCartProduct(s.ctx, productID, count, userID)
	require.Error(s.T(), err)
}

func (s *CartsStoreTestSuite) TestUpdatePaymentMethodSuccess() {
	userID := uint32(1)
	method := "credit_card"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `UPDATE users SET preferred_payment_method = \$1 WHERE id = \$2`

	s.dbMock.ExpectExec(query).WithArgs(method, userID).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err := s.store.UpdatePaymentMethod(s.ctx, userID, method)
	require.NoError(s.T(), err)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *CartsStoreTestSuite) TestUpdatePaymentMethodError() {
	userID := uint32(1)
	method := "credit_card"
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `UPDATE users SET preferred_payment_method = \$1 WHERE id = \$2`

	s.dbMock.ExpectExec(query).WithArgs(method, userID).WillReturnError(fmt.Errorf("some error"))

	err := s.store.UpdatePaymentMethod(s.ctx, userID, method)
	require.Error(s.T(), err)
}

func TestCartsStoreTestSuite(t *testing.T) {
	suite.Run(t, new(CartsStoreTestSuite))
}
