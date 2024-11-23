<<<<<<<< HEAD:internal/apps/order.go
package apps
========
package main_service
>>>>>>>> d5de27b ([HACK-2][improve] микросервис csat):internal/apps/main_service/order.go

import (
	"net/http"

	"github.com/gorilla/mux"
)

type orderDelivery interface {
	GetOrders(w http.ResponseWriter, r *http.Request)
	GetOrderByID(w http.ResponseWriter, r *http.Request)
	CreateOrderFromCart(w http.ResponseWriter, r *http.Request)
	GetNearestDeliveryDate(w http.ResponseWriter, r *http.Request)
}

type OrderApp struct {
	delivery orderDelivery
	router   *mux.Router
}

func NewOrderApp(r *mux.Router, delivery orderDelivery) OrderApp {
	return OrderApp{
		router:   r,
		delivery: delivery,
	}
}

func (o *OrderApp) InitOrderApp() *mux.Router {
	sub := o.router.Methods(http.MethodGet, http.MethodPost).Subrouter()
	sub.HandleFunc("/orders", o.delivery.GetOrders).Methods(http.MethodGet)
	sub.HandleFunc("/order/{id}", o.delivery.GetOrderByID).Methods(http.MethodGet)
	sub.HandleFunc("/orders", o.delivery.CreateOrderFromCart).Methods(http.MethodPost)
	sub.HandleFunc("/orders/nearest", o.delivery.GetNearestDeliveryDate).Methods(http.MethodGet)

	return sub
}
