package main_service

import (
	"github.com/gorilla/mux"
	"net/http"
)

type wishlistDelivery interface {
	AddProductToWishlists(w http.ResponseWriter, r *http.Request)
	CopyWishlist(w http.ResponseWriter, r *http.Request)
	RemoveFromWishlist(w http.ResponseWriter, r *http.Request)
	RenameWishlist(w http.ResponseWriter, r *http.Request)
	GetWishlistByLink(w http.ResponseWriter, r *http.Request)
	GetAllUserWishlists(w http.ResponseWriter, r *http.Request)
	CreateWishlist(w http.ResponseWriter, r *http.Request)
	DeleteWishlist(w http.ResponseWriter, r *http.Request)
}

type WishlistApp struct {
	delivery wishlistDelivery
	router   *mux.Router
}

func NewWishlistApp(r *mux.Router, delivery wishlistDelivery) *WishlistApp {
	return &WishlistApp{
		router:   r,
		delivery: delivery,
	}
}

func (w *WishlistApp) InitWishlistRoutes() *mux.Router {
	sub := w.router.Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete).Subrouter()
	sub.HandleFunc("/api/v1/wishlist/{link}", w.delivery.GetWishlistByLink).Methods(http.MethodGet)

	sub.HandleFunc("/api/v1/wishlist", w.delivery.CopyWishlist).Methods(http.MethodPost)
	sub.HandleFunc("/api/v1/wishlist", w.delivery.RenameWishlist).Methods(http.MethodPatch)

	sub.HandleFunc("/api/v1/wishlists/product", w.delivery.RemoveFromWishlist).Methods(http.MethodDelete)
	sub.HandleFunc("/api/v1/wishlists/product", w.delivery.AddProductToWishlists).Methods(http.MethodPost)

	sub.HandleFunc("/api/v1/wishlists", w.delivery.GetAllUserWishlists).Methods(http.MethodGet)
	sub.HandleFunc("/api/v1/wishlists", w.delivery.CreateWishlist).Methods(http.MethodPost)
	sub.HandleFunc("/api/v1/wishlists", w.delivery.DeleteWishlist).Methods(http.MethodDelete)

	return sub
}
