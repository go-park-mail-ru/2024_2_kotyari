package notifications

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderState string

const (
	AwaitingPayment            OrderState = "awaiting_payment"
	Paid                       OrderState = "paid"
	Delivered                  OrderState = "delivered"
	DefaultStateSwitchInterval            = "2 minutes"
)

type NotificationsStore struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewNotificationsStore(pool *pgxpool.Pool, logger *slog.Logger) *NotificationsStore {
	return &NotificationsStore{
		db:  pool,
		log: logger,
	}
}
