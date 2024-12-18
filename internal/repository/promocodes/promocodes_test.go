package promocodes

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
)

type PromoCodesTestSuite struct {
	suite.Suite
	dbMock pgxmock.PgxConnIface
	store  *PromoCodesStore
	ctx    context.Context
	logger *slog.Logger
}

func (s *PromoCodesTestSuite) SetupSuite() {
	var err error
	s.dbMock, err = pgxmock.NewConn()
	require.NoError(s.T(), err)
	s.logger = logger.InitLogger()
	s.store = &PromoCodesStore{
		db:  s.dbMock,
		log: s.logger,
	}
	requestID := "test-request-id"
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)
}

func (s *PromoCodesTestSuite) TearDownSuite() {
	s.dbMock.Close(s.ctx)
}

func (s *PromoCodesTestSuite) TestAddPromoCodeSuccess() {
	userID := uint32(123)
	promoID := uint32(456)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)
	const query = `
		insert into user_promocodes\(user_id, promo_id\)
		values \(\$1, \$2\);
	`

	s.dbMock.ExpectExec(query).
		WithArgs(userID, promoID).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := s.store.AddPromoCode(s.ctx, userID, promoID)
	require.NoError(s.T(), err)
}

func (s *PromoCodesTestSuite) TestAddPromoCodeQueryError() {
	userID := uint32(123)
	promoID := uint32(456)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `
		insert into user_promocodes\(user_id, promo_id\)
		values \(\$1, \$2\);
	`

	s.dbMock.ExpectExec(query).
		WithArgs(userID, promoID).
		WillReturnError(errors.New("query error"))

	err := s.store.AddPromoCode(s.ctx, userID, promoID)
	require.Error(s.T(), err)
	require.Contains(s.T(), err.Error(), "query error")
}

func (s *PromoCodesTestSuite) TestDeletePromoCodeSuccess() {
	userID := uint32(123)
	promoID := uint32(456)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `
		delete from user_promocodes
		where user_id = \$1 and promo_id = \$2;
	`

	s.dbMock.ExpectExec(query).
		WithArgs(userID, promoID).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err := s.store.DeletePromoCode(s.ctx, userID, promoID)
	require.NoError(s.T(), err)
}

func (s *PromoCodesTestSuite) TestDeletePromoCodeNoRows() {
	userID := uint32(123)
	promoID := uint32(456)
	requestID := uuid.New()
	s.ctx = context.WithValue(context.Background(), utils.RequestIDName, requestID)

	const query = `
		delete from user_promocodes
		where user_id = \$1 and promo_id = \$2;
	`

	s.dbMock.ExpectExec(query).
		WithArgs(userID, promoID).
		WillReturnResult(pgxmock.NewResult("DELETE", 0))

	err := s.store.DeletePromoCode(s.ctx, userID, promoID)
	require.NoError(s.T(), err)
}

func (s *PromoCodesTestSuite) TestGetPromoCodeSuccess() {
	userID := uint32(123)
	promoCodeName := "SPECIAL10"

	const query = `
		select p.id, up.user_id, p.name, p.bonus, p.updated_at, p.created_at
		from promocodes p
		join user_promocodes up on p.id = up.promo_id
		where up.user_id = \$1 and p.name = \$2;
	`
	expectedUpdatedAt, _ := time.Parse("2006-01-02", "2024-06-12")
	expectedCreatedAt, _ := time.Parse("2006-01-02", "2024-06-10")
	expectedPromoCode := PromoCodesDTO{
		ID:        1,
		UserID:    userID,
		Name:      promoCodeName,
		Bonus:     uint32(10),
		UpdatedAt: expectedUpdatedAt,
		CreatedAt: expectedCreatedAt,
	}

	s.dbMock.ExpectQuery(query).
		WithArgs(userID, promoCodeName).
		WillReturnRows(pgxmock.NewRows([]string{"id", "user_id", "name", "bonus", "updated_at", "created_at"}).
			AddRow(expectedPromoCode.ID, userID, promoCodeName, expectedPromoCode.Bonus, expectedPromoCode.UpdatedAt, expectedPromoCode.CreatedAt))

	promoCode, err := s.store.GetPromoCode(s.ctx, userID, promoCodeName)
	require.NoError(s.T(), err)
	require.Equal(s.T(), promoCodeName, promoCode.Name)
	require.Equal(s.T(), uint32(10), promoCode.Bonus)
}

func (s *PromoCodesTestSuite) TestGetPromoCodeNoRows() {
	userID := uint32(123)
	promoCodeName := "SPECIAL10"

	const query = `
		select p.id, up.user_id, p.name, p.bonus, p.updated_at, p.created_at
		from promocodes p
		join user_promocodes up on p.id = up.promo_id
		where up.user_id = \$1 and p.name = \$2;
	`

	s.dbMock.ExpectQuery(query).
		WithArgs(userID, promoCodeName).
		WillReturnError(errs.NoPromoCode)

	_, err := s.store.GetPromoCode(s.ctx, userID, promoCodeName)
	require.ErrorIs(s.T(), err, errs.NoPromoCode)
}

func TestPromoCodesTestSuite(t *testing.T) {
	suite.Run(t, new(PromoCodesTestSuite))
}
