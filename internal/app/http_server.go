package app

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/redis"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	sessions3 "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/sessions"
	userDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/handlers"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/sessions"
	userRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	sessions2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions"
	userServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	r        *mux.Router
	sessions *sessions3.SessionDelivery
	auth     *userDeliveryLib.UsersDelivery
	catalog  *handlers.CardsApp
	cfg      config
	log      *slog.Logger
}

func NewServer() *Server {

	redisClient := redis.MustLoadRedisClient()
	sessionRepo := sessions.NewSessionRepo(redisClient)
	sessionService := sessions2.NewSessionService(sessionRepo)
	sessionDelivery := sessions3.NewSessionDelivery(sessionRepo)

	dbPool := postgres.MustLoadPgxPool()
	userRepo := userRepoLib.NewUserRepo(dbPool)
	userService := userServiceLib.NewUserService(userRepo, *sessionService)
	userHandler := userDeliveryLib.NewUsersHandler(userService)
	//sessions := auth.NewSessions()
	//authManager := auth.NewAuthManager(sessions)
	return &Server{
		r:        mux.NewRouter(),
		auth:     userHandler,
		catalog:  handlers.NewCardsApp(db.NewProducts()),
		cfg:      initServer(),
		log:      logger.InitLogger(),
		sessions: sessionDelivery,
	}
}

func (s *Server) setupRoutes() {

	s.r.HandleFunc("/login", s.auth.GetUserByEmail).Methods(http.MethodPost)
	s.r.HandleFunc("/logout", s.sessions.Delete).Methods(http.MethodPost)
	s.r.HandleFunc("/signup", s.auth.CreateUser).Methods(http.MethodPost)
	s.r.HandleFunc("/catalog/products", s.catalog.Products).Methods(http.MethodGet)
	s.r.HandleFunc("/catalog/product/{id}", s.catalog.ProductByID).Methods(http.MethodGet)
	s.r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	//getUnimplemented := s.r.Methods(http.MethodGet).Subrouter()
	//getUnimplemented.HandleFunc("/basket", s.auth.UnimplementedRoutesHandler)
	//getUnimplemented.HandleFunc("/records", s.auth.UnimplementedRoutesHandler)
	//getUnimplemented.HandleFunc("/favorite", s.auth.UnimplementedRoutesHandler)
	//getUnimplemented.HandleFunc("/account", s.auth.UnimplementedRoutesHandler)
	//getUnimplemented.Use(middlewares.AuthMiddleware(s.sessions))
}

func (s *Server) Run() error {
	s.setupRoutes()

	handler := middlewares.CorsMiddleware(s.r, s.cfg.SessionLifetime)

	s.log.Info("starting server", slog.String("address:", s.cfg.ServerAddress))
	//log.Printf("Сервер запущен на: %s\n", s.cfg.ServerAddress)
	return http.ListenAndServe(s.cfg.ServerAddress, handler)
}
