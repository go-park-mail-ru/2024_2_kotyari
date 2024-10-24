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
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/handlers"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares"
	sessionsRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/sessions"
	userRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	sessionsServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions"
	userServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	r        *mux.Router
	sessions *sessionsDeliveryLib.SessionDelivery
	auth     *userDeliveryLib.UsersDelivery
	catalog  *handlers.CardsApp
	cfg      config
	log      *slog.Logger
}

func NewServer() *Server {

	redisClient := redis.MustLoadRedisClient()
	sessionsRepo := sessionsRepoLib.NewSessionRepo(redisClient)
	sessionsService := sessionsServiceLib.NewSessionService(sessionsRepo)
	sessionsDelivery := sessionsDeliveryLib.NewSessionDelivery(sessionsRepo)

	dbPool := postgres.MustLoadPgxPool()
	userRepo := userRepoLib.NewUserRepo(dbPool)
	userService := userServiceLib.NewUserService(userRepo, *sessionsService)
	userHandler := userDeliveryLib.NewUsersHandler(userService)

	//sessions := auth.NewSessions()
	//authManager := auth.NewAuthManager(sessions)
	return &Server{
		r:        mux.NewRouter(),
		auth:     userHandler,
		catalog:  handlers.NewCardsApp(db.NewProducts()),
		cfg:      initServer(),
		log:      logger.InitLogger(),
		sessions: sessionsDelivery,
	}
}

func (s *Server) setupRoutes() {

	s.r.HandleFunc("/login", s.auth.GetUserByEmail).Methods(http.MethodPost)
	s.r.HandleFunc("/logout", s.sessions.Delete).Methods(http.MethodPost)
	s.r.HandleFunc("/signup", s.auth.CreateUser).Methods(http.MethodPost)
	s.r.HandleFunc("/catalog", s.catalog.Products).Methods(http.MethodGet)
	s.r.HandleFunc("/product/{id}", s.catalog.ProductByID).Methods(http.MethodGet)
	s.r.HandleFunc("/", s.auth.GetUserById).Methods(http.MethodGet)
	s.r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	getUnimplemented := s.r.Methods(http.MethodGet).Subrouter()
	getUnimplemented.HandleFunc("/basket", func(w http.ResponseWriter, r *http.Request) {

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
