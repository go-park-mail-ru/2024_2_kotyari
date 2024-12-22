package rorders

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type DBConn interface {
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type OrdersRepo struct {
	db     pool.DBPool
	logger *slog.Logger
}

func NewOrdersRepo(db pool.DBPool, logger *slog.Logger) *OrdersRepo {
	return &OrdersRepo{
		db:     db,
		logger: logger,
	}
}
