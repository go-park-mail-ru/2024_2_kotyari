package app

import (
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/auth"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/handlers"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	r       *mux.Router
	auth    *auth.Manager
	catalog *handlers.CardsApp
	cfg     config
}

func NewServer() *Server {
	return &Server{
		r:       mux.NewRouter(),
		auth:    auth.NewAuthManager(auth.NewSessions()),
		catalog: handlers.NewCardsApp(db.NewProducts()),
		cfg:     initServer(),
	}
}

func (s *Server) setupRoutes() {
	s.r.HandleFunc("/login", s.auth.Login).Methods(http.MethodPost)
	s.r.HandleFunc("/logout", s.auth.Logout).Methods(http.MethodPost)
	s.r.HandleFunc("/signup", s.auth.SignUp).Methods(http.MethodPost)
	s.r.HandleFunc("/catalog/products", s.catalog.Products).Methods(http.MethodGet)
	s.r.HandleFunc("/catalog/product/{id}", s.catalog.ProductByID).Methods(http.MethodGet)
	s.r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	getUnimplemented := s.r.Methods(http.MethodGet).Subrouter()
	getUnimplemented.HandleFunc("/basket", s.auth.Soon)
	getUnimplemented.HandleFunc("/records", s.auth.Soon)
	getUnimplemented.HandleFunc("/favorite", s.auth.Soon)
	getUnimplemented.HandleFunc("/account", s.auth.Soon)
	getUnimplemented.Use(middlewares.AuthMiddleware(s.auth))
}

func (s *Server) Run() error {
	s.setupRoutes()

	// Оборачиваем маршруты CORS middleware
	handler := middlewares.CorsMiddleware(s.r, s.cfg.SessionLifetime)

	log.Printf("Сервер запущен на: %s\n", s.cfg.ServerAddress)
	return http.ListenAndServe(s.cfg.ServerAddress, handler)
}
