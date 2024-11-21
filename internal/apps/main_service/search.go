<<<<<<<< HEAD:internal/apps/main_service/search.go
package main_service
========
package go_main
>>>>>>>> bffcdd5 ([OZON-126][improve] микросервис авторизации):internal/app/go_main/search.go

import (
	"github.com/gorilla/mux"
	"net/http"
)

type searchDelivery interface {
	GetSearchTitleSuggestions(w http.ResponseWriter, r *http.Request)
	ProductSuggestions(w http.ResponseWriter, r *http.Request)
}
type SearchApp struct {
	delivery searchDelivery
	router   *mux.Router
}

func NewSearchApp(router *mux.Router, delivery searchDelivery) SearchApp {
	return SearchApp{
		delivery: delivery,
		router:   router,
	}
}
func (s *SearchApp) InitSearchRoutes() *mux.Router {
	sub := s.router.Methods(http.MethodGet).Subrouter()
	sub.HandleFunc("/search", s.delivery.GetSearchTitleSuggestions)
	sub.HandleFunc("/search/catalog", s.delivery.ProductSuggestions)
	return sub
}
