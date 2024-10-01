package app

import (
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/config"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	r       *mux.Router
	auth    *handlers.AuthApp
	catalog *handlers.CardsApp
	cfg     config.Server
}

func NewServer() *Server {
	return &Server{
		r:       mux.NewRouter(),
		auth:    handlers.NewApp(),
		catalog: handlers.NewCardsApp(db.NewProducts()),
		cfg:     config.InitServer(),
	}
}

func (s *Server) setupRoutes() {
	s.r.HandleFunc("/login", s.auth.Login).Methods(http.MethodPost)
	s.r.HandleFunc("/logout", s.auth.Logout).Methods(http.MethodPost)
	s.r.HandleFunc("/signup", s.auth.SignUp).Methods(http.MethodPost)
	s.r.HandleFunc("/catalog/products", s.catalog.Products).Methods(http.MethodGet)
	s.r.HandleFunc("/catalog/product/{id}", s.catalog.ProductByID).Methods(http.MethodGet)
	s.r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	s.r.HandleFunc("/basket", s.auth.IsLogin).Methods(http.MethodGet)
	s.r.HandleFunc("/records", s.auth.IsLogin).Methods(http.MethodGet)
	s.r.HandleFunc("/favorite", s.auth.IsLogin).Methods(http.MethodGet)
	s.r.HandleFunc("/account", s.auth.IsLogin).Methods(http.MethodGet)
	s.r.HandleFunc("/", s.auth.IsLogin).Methods(http.MethodGet)
}

func (s *Server) Run() error {
	s.setupRoutes()

	// Оборачиваем маршруты CORS middleware
	handler := corsMiddleware(s.r)

	log.Printf("Сервер запущен на: %s\n", s.cfg.ServerAddress)
	return http.ListenAndServe(s.cfg.ServerAddress, handler)
}
