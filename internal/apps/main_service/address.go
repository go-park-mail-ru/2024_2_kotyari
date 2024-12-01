package main_service

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/gorilla/mux"
)

type addressesDelivery interface {
	UpdateAddressData(w http.ResponseWriter, r *http.Request)
	GetAddress(w http.ResponseWriter, r *http.Request)
	GetAddressSuggestions(w http.ResponseWriter, r *http.Request)
}

type AddressApp struct {
	delivery    addressesDelivery
	errResolver errs.GetErrorCode
	log         *slog.Logger
	router      *mux.Router
}

func NewAddressApp(delivery addressesDelivery, errResolver errs.GetErrorCode, log *slog.Logger, router *mux.Router) AddressApp {
	return AddressApp{
		delivery:    delivery,
		errResolver: errResolver,
		log:         log,
		router:      router,
	}
}

func (a *AddressApp) InitRoutes() *mux.Router {
	sub := a.router.Methods(http.MethodGet, http.MethodPut).Subrouter()
	sub.HandleFunc("/address", a.delivery.GetAddress).Methods(http.MethodGet)
	sub.HandleFunc("/search/address", a.delivery.GetAddressSuggestions).Methods(http.MethodGet)
	sub.HandleFunc("/address", a.delivery.UpdateAddressData).Methods(http.MethodPut)

	return sub
}
