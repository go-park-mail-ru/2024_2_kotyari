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
	r        *mux.Router
	sessions auth.SessionInterface
	auth     auth.AuthManager
	catalog  *handlers.CardsApp
	cfg      config
}

func NewServer() *Server {
	// Нужно для middleware, потому что там отдельно необходимы сессии
	sessions := auth.NewSessions()
	authManager := auth.NewAuthManager(sessions)
	return &Server{
		r:        mux.NewRouter(),
		sessions: sessions,
		auth:     authManager,
		catalog:  handlers.NewCardsApp(db.NewProducts()),
		cfg:      initServer(),
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
	getUnimplemented.HandleFunc("/basket", s.auth.UnimplementedRoutesHandler)
	getUnimplemented.HandleFunc("/records", s.auth.UnimplementedRoutesHandler)
	getUnimplemented.HandleFunc("/favorite", s.auth.UnimplementedRoutesHandler)
	getUnimplemented.HandleFunc("/account", s.auth.UnimplementedRoutesHandler)
	getUnimplemented.Use(middlewares.AuthMiddleware(s.sessions))
}

func (s *Server) Run() error {
	s.setupRoutes()

	// Оборачиваем маршруты CORS middleware
	handler := middlewares.CorsMiddleware(s.r, s.cfg.SessionLifetime)

	log.Printf("Сервер запущен на: %s\n", s.cfg.ServerAddress)
	return http.ListenAndServe(s.cfg.ServerAddress, handler)
}
