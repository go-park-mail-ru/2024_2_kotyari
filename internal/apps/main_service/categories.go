<<<<<<<< HEAD:internal/apps/main_service/categories.go
package main_service
========
package go_main
>>>>>>>> bffcdd5 ([OZON-126][improve] микросервис авторизации):internal/app/go_main/categories.go

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
