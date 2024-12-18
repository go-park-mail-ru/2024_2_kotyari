package product

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
	"regexp"
	"testing"
)

type ProductsStoreTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *ProductsStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *ProductsStoreTestSuite) SetupTest() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &ProductsStore{db: s.dbMock, log: s.logger}
	s.ctx = context.Background()
}

func (s *ProductsStoreTestSuite) TearDownTest() {
	s.dbMock.Close(s.ctx)
}

func (s *ProductsStoreTestSuite) TestGetAllProductsSuccess() {
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `
		SELECT p.id, p.title, p.price, p.original_price, p.rating,
		       p.discount, p.image_url, p.description
		FROM products p
			where p.active = true and p.count > 0
		ORDER BY p.created_at DESC;
	`

	rows := pgxmock.NewRows([]string{
		"id", "title", "price", "original_price", "rating", "discount", "image_url", "description",
	}).
		AddRow(uint32(1), "Product 1", uint32(100), uint32(120), float32(4.5), uint32(10), "image1.png", "Description 1").
		AddRow(uint32(2), "Product 2", uint32(200), uint32(250), float32(4.8), uint32(20), "image2.png", "Description 2")

	s.dbMock.ExpectQuery(query).WillReturnRows(rows)

	products, err := s.store.GetAllProducts(s.ctx)
	require.NoError(s.T(), err)
	require.Len(s.T(), products, 2)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

//func (s *ProductsStoreTestSuite) TestGetProductByIDSuccess() {
//	productID := uint64(1)
//	requestID := uuid.New()
//	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)
//
//	queryProductInfo := `
//  SELECT
//      p.id, p.title, p.count,
//      p.price, p.original_price, p.discount,
//      p.rating,  p.description, p.characteristics::jsonb,
//      s.id, s.name, s.logo
//  FROM products p
//      JOIN sellers s ON p.seller_id = s.id
//  WHERE p.id = $1 AND p.active = true;`
//
//	const queryGetProductOptions = `
//  SELECT id, name, value
//  FROM product_options
//  WHERE product_id = $1;
//`
//	const queryGetProductImages = `
//  SELECT url
//  FROM product_images
//  WHERE product_id = $1;
//`
//	//	characteristics := `{"color": "red", "size": "L"}`
//
//	productRow := pgxmock.NewRows([]string{
//		"id", "title", "count", "price", "original_price", "discount",
//		"rating", "description", "characteristics", "seller_id", "seller_name", "seller_logo",
//	}).
//		AddRow(productID, "Test Product", 10, 100, 120, 20, 4.5, "Description", `{}`, 1, "Seller 1", "logo.png")
//
//	s.dbMock.ExpectQuery(regexp.QuoteMeta(queryProductInfo)).WithArgs(productID).WillReturnRows(productRow)
//
//	s.dbMock.ExpectQuery(queryGetProductCategories).WithArgs(productID).
//		WillReturnRows(pgxmock.NewRows([]string{"id", "name", "picture"}).AddRow(1, "Category 1", "pic.png"))
//
//	s.dbMock.ExpectQuery(queryGetProductOptions).WithArgs(productID).
//		WillReturnRows(pgxmock.NewRows([]string{"id", "name", "value"}))
//
//	s.dbMock.ExpectQuery(queryGetProductImages).WithArgs(productID).
//		WillReturnRows(pgxmock.NewRows([]string{"url"}).AddRow("image1.png"))
//
//	product, err := s.store.GetProductByID(s.ctx, productID)
//	require.NoError(s.T(), err)
//	require.Equal(s.T(), product.ID, productID)
//	require.Len(s.T(), product.Categories, 1)
//	require.Len(s.T(), product.Images, 1)
//
//	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
//}

func (s *ProductsStoreTestSuite) TestGetProductCountSuccess() {
	productID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select count from products where id = \$1;`

	row := pgxmock.NewRows([]string{"count"}).AddRow(10)

	s.dbMock.ExpectQuery(query).WithArgs(productID).WillReturnRows(row)

	count, err := s.store.GetProductCount(s.ctx, productID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), uint32(0), count)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestGetProductCategoriesSuccess() {
	productID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `
    SELECT 
        c.id, c.name,c.picture, c.link_to
    FROM product_categories pc
        JOIN categories c ON pc.category_id = c.id
    WHERE pc.product_id = \$1 AND pc.active = true;
	`

	rows := pgxmock.NewRows([]string{"id", "name", "picture", "link_to"}).
		AddRow(uint32(1), "Category 1", "pic1.png", "Link to 1").
		AddRow(uint32(2), "Category 2", "pic2.png", "Link to 2")

	s.dbMock.ExpectQuery(query).WithArgs(productID).WillReturnRows(rows)

	categories, err := s.store.GetProductCategories(s.ctx, productID)
	require.NoError(s.T(), err)
	require.Len(s.T(), categories, 2)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestGetProductImagesSuccess() {
	productID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	queryGetImagesProduct := `
		SELECT image_url
		FROM product_images
		WHERE product_id = \$1;
	`

	rows := pgxmock.NewRows([]string{"url"}).
		AddRow("image1.png").
		AddRow("image2.png")

	s.dbMock.ExpectQuery(queryGetImagesProduct).WithArgs(productID).WillReturnRows(rows)

	images, err := s.store.getProductImages(s.ctx, uint32(productID))

	require.NoError(s.T(), err)
	require.Len(s.T(), images, 2)
	require.Equal(s.T(), images[0].Url, "image1.png")
	require.Equal(s.T(), images[1].Url, "image2.png")

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestGetProductImagesNoResults() {
	productID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	queryGetImagesProduct := `
		SELECT image_url
		FROM product_images
		WHERE product_id = \$1;
	`

	s.dbMock.ExpectQuery(queryGetImagesProduct).WithArgs(productID).WillReturnRows(pgxmock.NewRows([]string{"url"}))

	images, err := s.store.getProductImages(s.ctx, uint32(productID))

	require.Error(s.T(), err)
	require.Nil(s.T(), images)
	require.Equal(s.T(), err, errs.ImagesDoesNotExists)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestGetProductImagesQueryError() {
	productID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	s.dbMock.ExpectQuery(regexp.QuoteMeta(queryGetImagesProduct)).WithArgs(productID).WillReturnError(fmt.Errorf("query execution error"))

	images, err := s.store.getProductImages(s.ctx, uint32(productID))

	require.Error(s.T(), err)
	require.Nil(s.T(), images)
	require.Equal(s.T(), err.Error(), "query execution error")

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestGetProductOptionsSuccess() {
	productID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	queryGetOptions := `
		SELECT po.values::jsonb FROM product_options po WHERE po.product_id = $1;
	`

	optionValuesJSON := `
	[
		{
			"title": "Color",
			"type": "select",
			"options": [
				{"link": "red", "value": "Red"}
			]
		},
		{
			"title": "Size",
			"type": "select",
			"options": [
				{"link": "l", "value": "L"}
			]
		}
	]`

	rows := pgxmock.NewRows([]string{"values"}).AddRow([]byte(optionValuesJSON))

	s.dbMock.ExpectQuery(regexp.QuoteMeta(queryGetOptions)).WithArgs(productID).WillReturnRows(rows)

	options, err := s.store.getProductOptions(s.ctx, uint32(productID))

	require.NoError(s.T(), err)
	require.Len(s.T(), options.Values, 2)

	require.Equal(s.T(), options.Values[0].Title, "Color")
	require.Equal(s.T(), options.Values[0].Type, "select")
	require.Len(s.T(), options.Values[0].Options, 1)
	require.Equal(s.T(), options.Values[0].Options[0].Link, "red")
	require.Equal(s.T(), options.Values[0].Options[0].Value, "Red")

	require.Equal(s.T(), options.Values[1].Title, "Size")
	require.Equal(s.T(), options.Values[1].Type, "select")
	require.Len(s.T(), options.Values[1].Options, 1)
	require.Equal(s.T(), options.Values[1].Options[0].Link, "l")
	require.Equal(s.T(), options.Values[1].Options[0].Value, "L")

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestGetProductOptionsNoResults() {
	productID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	queryGetOptions := `
		SELECT po.values::jsonb FROM product_options po WHERE po.product_id = $1;
	`

	s.dbMock.ExpectQuery(regexp.QuoteMeta(queryGetOptions)).WithArgs(productID).WillReturnRows(pgxmock.NewRows([]string{"option_values"}))

	options, err := s.store.getProductOptions(s.ctx, uint32(productID))

	require.Error(s.T(), err)
	require.Empty(s.T(), options.Values)
	require.Equal(s.T(), err, errs.OptionsDoesNotExists)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestGetProductOptionsQueryError() {
	productID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	s.dbMock.ExpectQuery(regexp.QuoteMeta(queryGetOptions)).WithArgs(productID).WillReturnError(fmt.Errorf("query execution error"))

	options, err := s.store.getProductOptions(s.ctx, uint32(productID))

	require.Error(s.T(), err)
	require.Empty(s.T(), options.Values)
	require.Equal(s.T(), err.Error(), "query execution error")

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestUpdateProductRatingSuccess() {
	productID := uint32(1)
	newRating := float32(4.5)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `
		update products
		set rating = $2
		where id = $1;
	`

	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(productID, newRating).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err := s.store.UpdateProductRating(s.ctx, productID, newRating)

	require.NoError(s.T(), err)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestUpdateProductRatingNoRowsAffected() {
	productID := uint32(1)
	newRating := float32(4.5)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `
		update products
		set rating = $2
		where id = $1;
	`
	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(productID, newRating).
		WillReturnResult(pgxmock.NewResult("UPDATE", 0))

	err := s.store.UpdateProductRating(s.ctx, productID, newRating)

	require.Error(s.T(), err)
	require.Equal(s.T(), err, errs.ProductNotFound)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ProductsStoreTestSuite) TestUpdateProductRatingQueryError() {
	productID := uint32(1)
	newRating := float32(4.5)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `
		update products
		set rating = $2
		where id = $1;
	`

	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(productID, newRating).
		WillReturnError(fmt.Errorf("query execution error"))

	err := s.store.UpdateProductRating(s.ctx, productID, newRating)

	require.Error(s.T(), err)
	require.Equal(s.T(), err.Error(), "query execution error")

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func TestCartsStoreTestSuite(t *testing.T) {
	suite.Run(t, new(ProductsStoreTestSuite))
}
