package notifications

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/pool"
	"log/slog"
)

type OrderState string

const (
	AwaitingPayment            OrderState = "awaiting_payment"
	Paid                       OrderState = "paid"
	Delivered                  OrderState = "delivered"
	DefaultStateSwitchInterval            = "2 minutes"
)

type NotificationsStore struct {
	db  pool.DBPool
	log *slog.Logger
}

func NewNotificationsStore(pool pool.DBPool, logger *slog.Logger) *NotificationsStore {
	return &NotificationsStore{
		db:  pool,
		log: logger,
	}
}
