package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

type searchDelivery interface {
	GetSearchTitleSuggestions(w http.ResponseWriter, r *http.Request)
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

	return sub
}
