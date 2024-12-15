package notifications

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"log/slog"
)

type orderStatusesChanger interface {
	Run()
}

type NotificationsApp struct {
	orderStatusesChanger orderStatusesChanger
	log                  *slog.Logger
	config               configs.ServiceViperConfig
}

func NewNotificationsApp(config map[string]any, orderStatusesChanger orderStatusesChanger, logger *slog.Logger) *NotificationsApp {
	cfg := configs.ParseServiceViperConfig(config)

	return &NotificationsApp{
		orderStatusesChanger: orderStatusesChanger,
		log:                  logger,
		config:               cfg,
	}
}
