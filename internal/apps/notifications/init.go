package notifications

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/gorilla/mux"
	"net/http"
)

type notificationsDelivery interface {
	Listen(w http.ResponseWriter, r *http.Request)
}

type NotificationsApp struct {
	notificationsDelivery notificationsDelivery
	router                *mux.Router
	config                configs.ServiceViperConfig
}

func NewNotificationsApp(config map[string]any, notificationsDelivery notificationsDelivery, router *mux.Router) *NotificationsApp {
	cfg := configs.ParseServiceViperConfig(config)

	return &NotificationsApp{
		notificationsDelivery: notificationsDelivery,
		router:                router,
		config:                cfg,
	}
}
