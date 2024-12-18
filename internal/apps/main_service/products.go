package main_service

import (
	"net/http"

	"github.com/gorilla/mux"
)

type productsDelivery interface {
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
}

type recDelivery interface {
	GetRecommendations(w http.ResponseWriter, r *http.Request)
}

type ProductsApp struct {
	delivery    productsDelivery
	router      *mux.Router
	recDelivery recDelivery
}

func NewProductsApp(r *mux.Router, delivery productsDelivery, recDelivery recDelivery) *ProductsApp {
	return &ProductsApp{
		router:      r,
		delivery:    delivery,
		recDelivery: recDelivery,
	}
}

func (p *ProductsApp) InitProductsRoutes() *mux.Router {
	sub := p.router.Methods(http.MethodGet).Subrouter()

	sub.HandleFunc("/api/v1/catalog", p.delivery.GetAllProducts).Methods(http.MethodGet)
	sub.HandleFunc("/api/v1/product/{id}", p.delivery.GetProductById).Methods(http.MethodGet)
	sub.HandleFunc("/api/v1/product/{id}/recommendations", p.recDelivery.GetRecommendations).Methods(http.MethodGet)

	return sub
}
