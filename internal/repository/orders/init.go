package rorders

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type OrdersRepo struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewOrdersRepo(db *pgxpool.Pool, logger *slog.Logger) *OrdersRepo {
	return &OrdersRepo{
		db:     db,
		logger: logger,
	}
}
