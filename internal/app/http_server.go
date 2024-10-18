package app

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/auth"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/handlers"
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
	sessions := auth.NewSessions()
	authManager := auth.NewAuthManager(sessions)
	return &Server{
		r:       mux.NewRouter(),
		auth:    auth.NewAuthManager(auth.NewSessions()),
		catalog: handlers.NewCardsApp(db.NewProducts()),
		cfg:     initServer(),
		log:     logger.InitLogger(),
		r:        mux.NewRouter(),
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

	handler := middlewares.CorsMiddleware(s.r, s.cfg.SessionLifetime)

	s.log.Info("starting server", slog.String("address:", s.cfg.ServerAddress))
	//log.Printf("Сервер запущен на: %s\n", s.cfg.ServerAddress)
	return http.ListenAndServe(s.cfg.ServerAddress, handler)
}
