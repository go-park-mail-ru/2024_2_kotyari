<<<<<<<< HEAD:internal/apps/categories.go
package apps
========
package main_service
>>>>>>>> d5de27b ([HACK-2][improve] микросервис csat):internal/apps/main_service/categories.go

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
	ca.router.HandleFunc("/category/{link}", ca.delivery.GetProductsByCategoryLink).Methods("GET")
	ca.router.HandleFunc("/categories", ca.delivery.GetAllCategories).Methods("GET")
}