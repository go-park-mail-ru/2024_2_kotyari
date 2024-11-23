package apps

import (
	"net/http"

	"github.com/gorilla/mux"
)

type cartDelivery interface {
	ChangeCartProductQuantity(w http.ResponseWriter, r *http.Request)
	GetCart(w http.ResponseWriter, r *http.Request)
	AddProduct(w http.ResponseWriter, r *http.Request)
	RemoveProduct(w http.ResponseWriter, r *http.Request)
	ChangeCartProductSelectedState(w http.ResponseWriter, r *http.Request)
	ChangeAllCartProductsState(w http.ResponseWriter, r *http.Request)
	GetSelectedFromCart(w http.ResponseWriter, r *http.Request)
	UpdatePaymentMethod(w http.ResponseWriter, r *http.Request)
	RemoveSelected(w http.ResponseWriter, r *http.Request)
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
	sub := c.router.Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete).Subrouter()
	sub.HandleFunc("/cart", c.delivery.GetCart).Methods(http.MethodGet)
	sub.HandleFunc("/cart/product/{id}", c.delivery.ChangeCartProductQuantity).Methods(http.MethodPatch)
	sub.HandleFunc("/cart/product/{id}", c.delivery.AddProduct).Methods(http.MethodPost)
	sub.HandleFunc("/cart/product/{id}", c.delivery.RemoveProduct).Methods(http.MethodDelete)

	sub.HandleFunc("/cart/select/products", c.delivery.GetSelectedFromCart).Methods(http.MethodGet)
	sub.HandleFunc("/cart/select/product/{id}", c.delivery.ChangeCartProductSelectedState).Methods(http.MethodPatch)
	sub.HandleFunc("/cart/select/products", c.delivery.ChangeAllCartProductsState).Methods(http.MethodPatch)
	sub.HandleFunc("/cart/select/products", c.delivery.ChangeAllCartProductsState).Methods(http.MethodDelete)
	sub.HandleFunc("/cart/selected", c.delivery.RemoveSelected).Methods(http.MethodDelete)

	sub.HandleFunc("/cart/pay-method", c.delivery.UpdatePaymentMethod).Methods(http.MethodPatch)
	return sub
}
