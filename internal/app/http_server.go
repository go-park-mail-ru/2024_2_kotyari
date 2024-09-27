package app

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/config"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

type Server struct {
	CORS *cors.Cors
	r    *mux.Router
	a    *handlers.AuthApp
	c    *handlers.CardsApp
	cfg  config.Server
}

const (
	second = 1
	minute = 60 * second
	hour   = 60 * minute
)

func setUpCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080", "http://127.0.0.1:3000", "http://127.0.0.1:8080"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           hour,
		Debug:            true,
	})
}

func NewServer() *Server {
	return &Server{
		r:    mux.NewRouter(),
		a:    handlers.NewApp(),
		c:    handlers.NewCardsApp(db.NewProducts()),
		CORS: setUpCORS(),
		cfg:  config.InitServer(),
	}
}

func (s *Server) Run() {
	s.r.HandleFunc("/login", s.a.Login).Methods(http.MethodPost)
	s.r.HandleFunc("/logout", s.a.Logout).Methods(http.MethodPost)
	s.r.HandleFunc("/signup", s.a.SignUp).Methods(http.MethodPost)
	s.r.HandleFunc("/catalog/products", s.c.Products).Methods("GET")
	s.r.HandleFunc("/catalog/product/{id}", s.c.ProductByID).Methods("GET")
	s.r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	handler := s.CORS.Handler(s.r)

	log.Printf("Сервер запущен на: %s\n", s.cfg.ServerAddress)
	err := http.ListenAndServe(s.cfg.ServerAddress, handler)
	if err != nil {
		log.Fatal(err)
	}
}
