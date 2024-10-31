package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type cartDelivery interface {
	ChangeCartProductQuantity(w http.ResponseWriter, r *http.Request)
	GetCart(w http.ResponseWriter, r *http.Request)
}

type CartApp struct {
	delivery cartDelivery
	router   *mux.Router
}

func NewCartApp(r *mux.Router, delivery cartDelivery) CartApp {
	return CartApp{
		router:   r,
		delivery: delivery,
	}
}

func (c *CartApp) InitCartRoutes() *mux.Router {
	sub := c.router.Methods(http.MethodGet, http.MethodPost).Subrouter()
	sub.HandleFunc("/cart", c.delivery.GetCart).Methods(http.MethodGet)
	sub.HandleFunc("/cart/product/{id}", c.delivery.ChangeCartProductQuantity).Methods(http.MethodPost)
	return sub
}
