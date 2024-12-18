package main_service

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
	sub.HandleFunc("/api/v1/search", s.delivery.GetSearchTitleSuggestions)
	sub.HandleFunc("/api/v1/search/catalog", s.delivery.ProductSuggestions)
	return sub
}
