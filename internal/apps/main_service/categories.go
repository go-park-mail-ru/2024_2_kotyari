package main_service

import (
	"net/http"

	"github.com/gorilla/mux"
)

type categoriesDelivery interface {
	GetAllCategories(w http.ResponseWriter, r *http.Request)
	GetProductsByCategoryLink(w http.ResponseWriter, r *http.Request)
}

type CategoryApp struct {
	delivery categoriesDelivery
	router   *mux.Router
}

func NewCategoryApp(r *mux.Router, delivery categoriesDelivery) *CategoryApp {
	return &CategoryApp{
		router:   r,
		delivery: delivery,
	}
}

func (ca *CategoryApp) InitCategoriesRoutes() {
	ca.router.HandleFunc("/api/v1/category/{link}", ca.delivery.GetProductsByCategoryLink).Methods("GET")
	ca.router.HandleFunc("/api/v1/categories", ca.delivery.GetAllCategories).Methods("GET")
}
