package app

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/redis"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	sessionsDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/sessions"
	userDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	errResolveLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/handlers"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares"
	sessionsRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/sessions"
	userRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	sessionsServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions"
	userServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/gorilla/mux"
)

type usersDelivery interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

type sessionRemoverDelivery interface {
	Delete(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	r        *mux.Router
	sessions sessionRemoverDelivery
	auth     usersDelivery
	catalog  *handlers.CardsApp
	cfg      config
	log      *slog.Logger
}

func NewServer() (*Server, error) {
	errResolver := errResolveLib.NewErrorStore()
	redisClient, err := redis.LoadRedisClient()
	if err != nil {
		return nil, err
	}

	sessionsRepo := sessionsRepoLib.NewSessionRepo(redisClient)
	sessionsService := sessionsServiceLib.NewSessionService(sessionsRepo)
	sessionsDelivery := sessionsDeliveryLib.NewSessionDelivery(sessionsService, sessionsRepo, errResolver)

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		return nil, err
	}

	userRepo := userRepoLib.NewUsersStore(dbPool)
	userService := userServiceLib.NewUserService(userRepo, sessionsService)
	userHandler := userDeliveryLib.NewUsersHandler(userService, errResolver)

	return &Server{
		r:        mux.NewRouter(),
		auth:     userHandler,
		catalog:  handlers.NewCardsApp(db.NewProducts()),
		cfg:      initServer(),
		log:      logger.InitLogger(),
		sessions: sessionsDelivery,
	}, nil
}

func (s *Server) setupRoutes() {

	s.r.HandleFunc("/login", s.auth.LoginUser).Methods(http.MethodPost)
	s.r.HandleFunc("/logout", s.sessions.Delete).Methods(http.MethodPost)
	s.r.HandleFunc("/signup", s.auth.CreateUser).Methods(http.MethodPost)
	s.r.HandleFunc("/catalog", s.catalog.Products).Methods(http.MethodGet)
	s.r.HandleFunc("/product/{id}", s.catalog.ProductByID).Methods(http.MethodGet)
	s.r.HandleFunc("/", s.auth.GetUserById).Methods(http.MethodGet)

	getUnimplemented := s.r.Methods(http.MethodGet).Subrouter()
	getUnimplemented.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {

	})
	getUnimplemented.HandleFunc("/records", func(w http.ResponseWriter, r *http.Request) {

	})
	getUnimplemented.HandleFunc("/favorite", func(w http.ResponseWriter, r *http.Request) {

	})
	getUnimplemented.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {

	})
	getUnimplemented.Use(middlewares.AuthMiddleware)
}

func (s *Server) Run() error {
	s.setupRoutes()

	handler := middlewares.CorsMiddleware(s.r, s.cfg.SessionLifetime)

	s.log.Info("starting server", slog.String("address:", s.cfg.ServerAddress))
	return http.ListenAndServe(s.cfg.ServerAddress, handler)
}
