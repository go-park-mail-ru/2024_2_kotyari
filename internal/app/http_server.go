package app

import (
	"context"
	orders2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/orders"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/redis"
	addressDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/address"
	cartDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/cart"
	categoryDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/category"
	csrf2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/csrf"
	fileDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/file"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/orders"
	productDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/product"
	profileDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/profile"
	sessionsDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/sessions"
	userDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	errResolveLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	addressRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/address"
	cartRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/cart"
	categoryRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/category"
	fileRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/file"
	rorders "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/orders"
	productRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/product"
	profileRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/profile"
	sessionsRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/sessions"
	userRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	addressServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/address"
	cartServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/cart"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/csrf"
	fileServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/file"
	imageServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/image"
	profileServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/profile"
	sessionsServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions"
	userServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
)

type categoryApp interface {
	InitCategoriesRoutes()
}

type filesDelivery interface {
	GetImage(w http.ResponseWriter, r *http.Request)
}

type usersDelivery interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

type profilesDelivery interface {
	GetProfile(writer http.ResponseWriter, request *http.Request)
	UpdateProfileData(writer http.ResponseWriter, request *http.Request)
	UpdateProfileAvatar(writer http.ResponseWriter, request *http.Request)
}

type addressesDelivery interface {
	UpdateAddressData(writer http.ResponseWriter, request *http.Request)
	GetAddress(writer http.ResponseWriter, request *http.Request)
}

type SessionDelivery interface {
	Delete(w http.ResponseWriter, r *http.Request)
	Get(ctx context.Context, sessionID string) (model.Session, error)
}

type productsApp interface {
	InitProductsRoutes() *mux.Router
}

type Server struct {
	r        *mux.Router
	sessions SessionDelivery
	auth     usersDelivery
	product  productsApp
	category categoryApp
	profile  profilesDelivery
	address  addressesDelivery
	cfg      config
	log      *slog.Logger
	files    filesDelivery
	cart     CartApp
	order    OrderApp
	csrf     csrfDelivery
}

type csrfDelivery interface {
	GetCsrf(w http.ResponseWriter, r *http.Request)
}

func NewServer() (*Server, error) {
	log := logger.InitLogger()
	router := mux.NewRouter()

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		return nil, err
	}

	inputValidator := utils.NewInputValidator()

	errResolver := errResolveLib.NewErrorStore()

	redisClient, err := redis.LoadRedisClient()
	if err != nil {
		return nil, err
	}

	fileRepo, err := fileRepoLib.NewFilesRepo(log)
	if err != nil {
		return nil, err
	}
	fileService := fileServiceLib.NewFilesUsecase(fileRepo, log)
	imageService := imageServiceLib.NewImagesUsecase(fileService)

	sessionsRepo := sessionsRepoLib.NewSessionRepo(redisClient)
	sessionsService := sessionsServiceLib.NewSessionService(sessionsRepo)
	sessionsDelivery := sessionsDeliveryLib.NewSessionDelivery(sessionsRepo, errResolver)

	userRepo := userRepoLib.NewUsersStore(dbPool)
	userService := userServiceLib.NewUserService(userRepo, sessionsService)
	userHandler := userDeliveryLib.NewUsersHandler(userService, inputValidator, errResolver)

	categoryRepo := categoryRepoLib.NewCategoriesStore(dbPool, log)
	categoryDelivery := categoryDeliveryLib.NewCategoriesDelivery(categoryRepo, log, errResolver)

	profileRepo := profileRepoLib.NewProfileRepo(dbPool, log)
	profileService := profileServiceLib.NewProfileService(imageService, profileRepo, log)
	profileHandler := profileDeliveryLib.NewProfilesHandler(profileService, inputValidator, log)

	addressRepo := addressRepoLib.NewAddressRepo(dbPool, log)
	addressService := addressServiceLib.NewAddressService(addressRepo, log)
	addressHandler := addressDeliveryLib.NewAddressHandler(addressService, log)

	prodRepo := productRepoLib.NewProductsStore(dbPool, log)

	ca := NewCategoryApp(router, categoryDelivery)

	cartRepo := cartRepoLib.NewCartsStore(dbPool, log)
	cartService := cartServiceLib.NewCartManager(cartRepo, prodRepo, log)
	cartHandler := cartDeliveryLib.NewCartHandler(cartService, cartRepo, errResolver, log)
	cartApp := NewCartApp(router, cartHandler)

	prodHandler := productDeliveryLib.NewProductHandler(errResolver, prodRepo, log, cartRepo)
	pa := NewProductsApp(router, prodHandler)

	fileDelivery := fileDeliveryLib.NewFilesDelivery(fileRepo)

	ordersRepo := rorders.NewOrdersRepo(dbPool, log)
	ordersManager := orders2.NewOrdersManager(ordersRepo, log, cartRepo)
	ordersHandler := orders.NewOrdersHandler(ordersManager, log, errResolver)
	orderApp := NewOrderApp(router, ordersHandler)

	csrfUsecase := csrf.NewCscfUsecase()
	csrfHandler := csrf2.NewCsrfDelivery(csrfUsecase, sessionsDelivery)

	return &Server{
		r:        router,
		auth:     userHandler,
		product:  pa,
		profile:  profileHandler,
		address:  addressHandler,
		category: ca,
		cfg:      initServer(),
		log:      log,
		sessions: sessionsDelivery,
		files:    fileDelivery,
		cart:     cartApp,
		order:    orderApp,
		csrf:     csrfHandler,
	}, nil
}

func (s *Server) setupRoutes() {
	errResolver := errResolveLib.NewErrorStore()

	csrfMiddleware := middlewares.CSRFMiddleware(csrf.NewCscfUsecase(), s.sessions)

	subProd := s.product.InitProductsRoutes()
	subProd.Use(middlewares.AuthMiddleware(s.sessions, errResolver))
	subProd.Use(csrfMiddleware)

	s.category.InitCategoriesRoutes()

	subCart := s.cart.InitCartRoutes()
	subCart.Use(middlewares.AuthMiddleware(s.sessions, errResolver))
	subCart.Use(csrfMiddleware)

	subOrder := s.order.InitOrderApp()
	subOrder.Use(middlewares.AuthMiddleware(s.sessions, errResolver))
	subOrder.Use(csrfMiddleware)

	s.r.HandleFunc("/login", s.auth.LoginUser).Methods(http.MethodPost)
	s.r.HandleFunc("/logout", s.sessions.Delete).Methods(http.MethodPost)
	s.r.HandleFunc("/signup", s.auth.CreateUser).Methods(http.MethodPost)

	authSub := s.r.Methods(http.MethodGet, http.MethodPost, http.MethodPut).Subrouter()
	authSub.HandleFunc("/csrf", s.csrf.GetCsrf).Methods(http.MethodGet)
	authSub.Use(middlewares.AuthMiddleware(s.sessions, errResolver))

	s.r.HandleFunc("/files/{name}", s.files.GetImage).Methods(http.MethodGet)

	csrfProtected := authSub.Methods(http.MethodGet, http.MethodPost, http.MethodPut).Subrouter()

	csrfProtected.HandleFunc("/account", s.profile.GetProfile).Methods(http.MethodGet)
	csrfProtected.HandleFunc("/account", s.profile.UpdateProfileData).Methods(http.MethodPut)
	csrfProtected.HandleFunc("/account/avatar", s.profile.UpdateProfileAvatar).Methods(http.MethodPut)
	csrfProtected.HandleFunc("/address", s.address.GetAddress).Methods(http.MethodGet)
	csrfProtected.HandleFunc("/address", s.address.UpdateAddressData).Methods(http.MethodPut)
	csrfProtected.Use(csrfMiddleware)

	s.r.HandleFunc("/", s.auth.GetUserById).Methods(http.MethodGet)
}

func (s *Server) Run() error {
	s.setupRoutes()

	handler := middlewares.CorsMiddleware(s.r, s.cfg.SessionLifetime)

	s.log.Info("starting  server", slog.String("address:", s.cfg.ServerAddress))
	return http.ListenAndServe(s.cfg.ServerAddress, handler)
}
