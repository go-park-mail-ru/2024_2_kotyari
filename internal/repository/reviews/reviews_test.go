package reviews

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
	"regexp"
	"testing"
	"time"
)

type ReviewsStoreTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *ReviewsStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *ReviewsStoreTestSuite) SetupTest() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)

	s.logger = slog.Default()
	s.store = &ReviewsStore{db: s.dbMock, log: s.logger}
	s.ctx = context.Background()
}

func (s *ReviewsStoreTestSuite) TearDownTest() {
	s.dbMock.Close(s.ctx)
}

func (s *ReviewsStoreTestSuite) TestAddReviewSuccess() {
	productID := uint32(1)
	userID := uint32(2)
	review := model.Review{Text: "Great product!", Rating: 5, IsPrivate: false}
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `insert into reviews(product_id, user_id, text, rating, is_private) values ($1, $2, $3, $4, $5);`

	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(productID, userID, review.Text, review.Rating, review.IsPrivate).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := s.store.AddReview(s.ctx, productID, userID, review)

	require.NoError(s.T(), err)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestAddReviewFail() {
	productID := uint32(1)
	userID := uint32(2)
	review := model.Review{Text: "Great product!", Rating: 5, IsPrivate: false}
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `insert into reviews(product_id, user_id, text, rating, is_private) values ($1, $2, $3, $4, $5);`

	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(productID, userID, review.Text, review.Rating, review.IsPrivate).
		WillReturnError(errors.New("insert error"))

	err := s.store.AddReview(s.ctx, productID, userID, review)

	require.Error(s.T(), err)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestDeleteReviewSuccess() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `delete from reviews where product_id = $1 and user_id = $2;`

	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(productID, userID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err := s.store.DeleteReview(s.ctx, productID, userID)

	require.NoError(s.T(), err)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestDeleteReviewFail() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `delete from reviews where product_id = $1 and user_id = $2;`

	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(productID, userID).
		WillReturnError(errors.New("delete error"))

	err := s.store.DeleteReview(s.ctx, productID, userID)

	require.Error(s.T(), err)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestGetProductReviewsNoLoginSuccess() {
	productID := uint32(1)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	countQuery := `select count(id) from reviews where product_id = $1;`
	s.dbMock.ExpectQuery(regexp.QuoteMeta(countQuery)).WithArgs(productID).
		WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(2))

	query := `select r.text, r.rating, r.is_private, u.username, u.avatar_url, r.created_at 
		from reviews r join users u on u.id = r.user_id 
		where r.product_id = $1 order by r.created_at desc;`

	rows := pgxmock.NewRows([]string{"text", "rating", "is_private", "username", "avatar_url", "created_at"}).
		AddRow("Great product!", uint8(5), false, "john_doe", "avatar1.png", time.Now()).
		AddRow("Not bad", uint8(4), false, "jane_doe", "avatar2.png", time.Now())

	s.dbMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(productID).WillReturnRows(rows)

	reviews, err := s.store.GetProductReviewsNoLogin(s.ctx, productID, "date", "ASC")

	require.NoError(s.T(), err)
	require.Len(s.T(), reviews.Reviews, 2)
	require.Equal(s.T(), reviews.TotalReviewCount, uint32(0x0))
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestGetProductReviewsWithLoginSuccess() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	countQuery := `select count(id) from reviews where product_id = $1;`
	s.dbMock.ExpectQuery(regexp.QuoteMeta(countQuery)).WithArgs(productID).
		WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(3))

	userReviewQuery := `select r.text, r.rating, r.is_private, r.created_at from reviews r join users u ON u.id = r.user_id where r.product_id = $1 AND r.user_id = $2;`

	s.dbMock.ExpectQuery(regexp.QuoteMeta(userReviewQuery)).WithArgs(productID, userID).
		WillReturnRows(pgxmock.NewRows([]string{"text", "rating", "is_private", "created_at"}).
			AddRow("Awesome!", uint8(5), false, time.Now()))

	query := `select r.text, r.rating, r.is_private, u.username, u.avatar_url, r.created_at from reviews r join users u on u.id = r.user_id 
		where r.product_id = $1 and r.user_id != $2 order by r.created_at desc;`

	rows := pgxmock.NewRows([]string{"text", "rating", "is_private", "username", "avatar_url", "created_at"}).
		AddRow("Okay product", uint8(3), false, "jane_doe", "avatar2.png", time.Now())

	s.dbMock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(productID, userID).WillReturnRows(rows)

	reviews, err := s.store.GetProductReviewsWithLogin(s.ctx, productID, userID, "date", "ASC")

	require.NoError(s.T(), err)
	require.Equal(s.T(), reviews.TotalReviewCount, uint32(0x0))
	require.Equal(s.T(), reviews.UserReview.Text, "Awesome!")
	require.Len(s.T(), reviews.Reviews, 1)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestUpdateReviewSuccess() {
	productID := uint32(1)
	userID := uint32(2)
	review := model.Review{Text: "Updated review text", Rating: 4}
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `update reviews set text = $3, rating = $4 where product_id = $1 and user_id = $2;`

	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(productID, userID, review.Text, review.Rating).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err := s.store.UpdateReview(s.ctx, productID, userID, review)

	require.NoError(s.T(), err)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestUpdateReviewNoRowsAffected() {
	productID := uint32(1)
	userID := uint32(2)
	review := model.Review{Text: "Updated review text", Rating: 4}
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `update reviews set text = $3, rating = $4 where product_id = $1 and user_id = $2`

	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(productID, userID, review.Text, review.Rating).
		WillReturnResult(pgxmock.NewResult("UPDATE", 0)) // Симулируем "0 затронутых строк"

	err := s.store.UpdateReview(s.ctx, productID, userID, review)

	require.NoError(s.T(), err)

	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestUpdateReviewQueryError() {
	productID := uint32(1)
	userID := uint32(2)
	review := model.Review{Text: "Updated review text", Rating: 4}
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `update reviews set text = $3, rating = $4 where product_id = $1 and user_id = $2;`

	s.dbMock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(productID, userID, review.Text, review.Rating).
		WillReturnError(errors.New("update error"))

	err := s.store.UpdateReview(s.ctx, productID, userID, review)

	require.Error(s.T(), err)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestGetReviewNotFound() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select r.text, r.rating, r.is_private, r.created_at, u.username, u.avatar_url from reviews r join users u on u.id = r.user_id where product_id = $1 and user_id = $2;`

	s.dbMock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(productID, userID).
		WillReturnError(pgx.ErrNoRows)

	_, err := s.store.GetReview(s.ctx, productID, userID)

	require.ErrorIs(s.T(), err, errs.ReviewNotFound)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func (s *ReviewsStoreTestSuite) TestGetReviewQueryError() {
	productID := uint32(1)
	userID := uint32(2)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	query := `select r.text, r.rating, r.is_private, r.created_at, u.username, u.avatar_url from reviews r join users u on u.id = r.user_id where product_id = $1 and user_id = $2;`

	s.dbMock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(productID, userID).
		WillReturnError(errors.New("query error"))

	_, err := s.store.GetReview(s.ctx, productID, userID)

	require.Error(s.T(), err)
	require.NoError(s.T(), s.dbMock.ExpectationsWereMet())
}

func TestCartsStoreTestSuite(t *testing.T) {
	suite.Run(t, new(ReviewsStoreTestSuite))
}
