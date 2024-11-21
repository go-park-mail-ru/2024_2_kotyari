<<<<<<<< HEAD:internal/apps/main_service/products.go
package main_service
========
package go_main
>>>>>>>> bffcdd5 ([OZON-126][improve] микросервис авторизации):internal/app/go_main/products.go

import (
	"net/http"

	"github.com/gorilla/mux"
)

type productsDelivery interface {
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
}

type ProductsApp struct {
	delivery productsDelivery
	router   *mux.Router
}

func NewProductsApp(r *mux.Router, delivery productsDelivery) *ProductsApp {
	return &ProductsApp{
		router:   r,
		delivery: delivery,
	}
}

func (p *ProductsApp) InitProductsRoutes() *mux.Router {
	sub := p.router.Methods(http.MethodGet).Subrouter()

	sub.HandleFunc("/catalog", p.delivery.GetAllProducts).Methods(http.MethodGet)
	sub.HandleFunc("/product/{id}", p.delivery.GetProductById).Methods(http.MethodGet)

	return sub
}
