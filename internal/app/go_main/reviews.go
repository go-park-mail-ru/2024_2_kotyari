package go_main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type reviewsDelivery interface {
	GetProductReviews(w http.ResponseWriter, r *http.Request)
	AddReview(w http.ResponseWriter, r *http.Request)
	UpdateReview(w http.ResponseWriter, r *http.Request)
	DeleteReview(w http.ResponseWriter, r *http.Request)
}

type ReviewsApp struct {
	reviewsDelivery reviewsDelivery
	router          *mux.Router
}

func NewReviewsApp(router *mux.Router, delivery reviewsDelivery) ReviewsApp {
	return ReviewsApp{
		reviewsDelivery: delivery,
		router:          router,
	}
}

func (app *ReviewsApp) InitRoutes() *mux.Router {
	sub := app.router.Methods(http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete).Subrouter()
	sub.HandleFunc("/product/{id}/reviews", app.reviewsDelivery.GetProductReviews).Methods(http.MethodGet)
	sub.HandleFunc("/product/{id}/reviews", app.reviewsDelivery.AddReview).Methods(http.MethodPost)
	sub.HandleFunc("/product/{id}/reviews", app.reviewsDelivery.UpdateReview).Methods(http.MethodPut)
	sub.HandleFunc("/product/{id}/reviews", app.reviewsDelivery.DeleteReview).Methods(http.MethodDelete)

	return sub
}
