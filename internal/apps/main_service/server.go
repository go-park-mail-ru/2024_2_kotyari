package main_service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	profilegrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/profile/gen"
	usergrpc "github.com/go-park-mail-ru/2024_2_kotyari/api/protos/user/gen"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/redis"
	addressDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/address"
	cartDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/cart"
	categoryDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/category"
	csrfDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/csrf"
	fileDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/file"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/orders"
	productDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/product"
	profileDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/profile"
	promocodesDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/promocodes"
	reviewsDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/reviews"
	searchDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/search"
	sessionsDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/sessions"
	userDeliveryLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/delivery/user"
	errResolveLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	metrics "github.com/go-park-mail-ru/2024_2_kotyari/internal/metrics/http"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares"
	metricsMiddleware "github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares/metrics"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	addressRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/address"
	cartRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/cart"
	categoryRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/category"
	fileRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/file"
	rorders "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/orders"
	productRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/product"
	reviewsRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/reviews"
	searchRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/search"
	sessionsRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/sessions"
	addressServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/address"
	cartServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/cart"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/csrf"
	fileServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/file"
	imageServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/image"
	ordersServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/orders"
	productServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/product"
	reviewsServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/reviews"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	mainService          = "main_service"
	ratingUpdaterService = "rating_updater"
	profileService       = "profile_go"
	userService          = "user_go"
	promocodesService    = "promocodes_go"
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
	ChangePassword(writer http.ResponseWriter, request *http.Request)
}

type promoCodesDelivery interface {
	GetUserPromoCodes(w http.ResponseWriter, r *http.Request)
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
	r          *mux.Router
	sessions   SessionDelivery
	auth       usersDelivery
	product    productsApp
	category   categoryApp
	profile    profilesDelivery
	address    addressesDelivery
	cfg        configs.ServiceViperConfig
	log        *slog.Logger
	files      filesDelivery
	cart       CartApp
	order      OrderApp
	csrf       csrfDelivery
	reviews    ReviewsApp
	search     SearchApp
	promoCodes promoCodesDelivery
}

type csrfDelivery interface {
	GetCsrf(w http.ResponseWriter, r *http.Request)
}

func NewServer() (*Server, error) {

	log := logger.InitLogger()
	router := mux.NewRouter()
	v, err := configs.SetupViper()
	if err != nil {
		return nil, err
	}

	httpMetrics, err := metrics.NewHTTPMetrics("main")
	if err != nil {
		return nil, err
	}

	metricsM := metricsMiddleware.CreateMetricsMiddleware(httpMetrics)

	router.Use(metricsM)
	router.Use(middlewares.RequestIDMiddleware)

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

	sessionsRepo := sessionsRepoLib.NewSessionRepo(redisClient, log)
	sessionService := sessions.NewSessionService(sessionsRepo, log)
	sessionsDelivery := sessionsDeliveryLib.NewSessionDelivery(sessionsRepo, errResolver)

	userGRPCCfg := v.GetStringMap(userService)
	userCfg, err := configs.ParseServiceViperConfig(userGRPCCfg)
	if err != nil {
		return nil, err
	}

	userConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", userCfg.Domain, userCfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	userClient := usergrpc.NewUserServiceClient(userConn)

	userHandler := userDeliveryLib.NewUsersDelivery(userClient, inputValidator, sessionService, errResolver, log)

	categoryRepo := categoryRepoLib.NewCategoriesStore(dbPool, log)
	categoryDelivery := categoryDeliveryLib.NewCategoriesDelivery(categoryRepo, log, errResolver)

	addressRepo := addressRepoLib.NewAddressRepo(dbPool, log)
	addressService := addressServiceLib.NewAddressService(addressRepo, log)
	addressHandler := addressDeliveryLib.NewAddressHandler(addressService, errResolver, log)

	profileGRPCCfg := v.GetStringMap(profileService)
	profileCfg, err := configs.ParseServiceViperConfig(profileGRPCCfg)
	if err != nil {
		return nil, err
	}

	profileConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", profileCfg.Domain, profileCfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	profileClient := profilegrpc.NewProfileClient(profileConn)

	profileHandler := profileDeliveryLib.NewProfilesHandler(profileClient, log, addressRepo, imageService, errResolver)
	prodRepo := productRepoLib.NewProductsStore(dbPool, log)
	ca := NewCategoryApp(router, categoryDelivery)

	promoCodesGRPC, err := promocodesDeliveryLib.NewPromoCodesGRPC(v.GetStringMap(promocodesService), errResolver, log)
	if err != nil {
		return nil, err
	}

	cartRepo := cartRepoLib.NewCartsStore(dbPool, log)
	cartService := cartServiceLib.NewCartManager(cartRepo, promoCodesGRPC, prodRepo, log)
	cartHandler := cartDeliveryLib.NewCartHandler(cartService, cartRepo, errResolver, log)
	cartApp := NewCartApp(router, cartHandler)

	productService := productServiceLib.NewProductService(cartRepo, prodRepo, log)
	prodHandler := productDeliveryLib.NewProductHandler(errResolver, prodRepo, productService, log, cartRepo)
	pa := NewProductsApp(router, prodHandler)

	fileDelivery := fileDeliveryLib.NewFilesDelivery(fileRepo)

	ordersRepo := rorders.NewOrdersRepo(dbPool, log)
	ordersManager := ordersServiceLib.NewOrdersManager(ordersRepo, promoCodesGRPC, log, cartRepo)
	ordersHandler := orders.NewOrdersHandler(ordersManager, log, errResolver)
	orderApp := NewOrderApp(router, ordersHandler)

	csrfUsecase := csrf.NewCscfUsecase()
	csrfHandler := csrfDeliveryLib.NewCsrfDelivery(csrfUsecase, sessionsDelivery)

	reviewsGRPCCfg := v.GetStringMap(ratingUpdaterService)
	reviewsGRPC, err := reviewsDeliveryLib.NewRatingUpdaterGRPC(reviewsGRPCCfg, log)
	if err != nil {
		return nil, err
	}

	reviewsRepo := reviewsRepoLib.NewReviewsStore(dbPool, log)
	reviewsManager, err := reviewsServiceLib.NewReviewsService(reviewsRepo, inputValidator, log, reviewsGRPC)
	if err != nil {
		return nil, err
	}

	reviewsHandler := reviewsDeliveryLib.NewReviewsHandler(reviewsManager, inputValidator, errResolver, log)
	reviewsApp := NewReviewsApp(router, reviewsHandler)

	searchRepo := searchRepoLib.NewSearchStore(dbPool, log)
	searchHandler := searchDeliveryLib.NewSearchDelivery(searchRepo, errResolver, log)
	searchApp := NewSearchApp(router, searchHandler)

	cfg := v.GetStringMap(mainService)
	mainCfg, err := configs.ParseServiceViperConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &Server{
		r:          router,
		auth:       userHandler,
		product:    pa,
		profile:    profileHandler,
		address:    addressHandler,
		category:   ca,
		cfg:        mainCfg,
		log:        log,
		sessions:   sessionsDelivery,
		files:      fileDelivery,
		cart:       cartApp,
		order:      orderApp,
		csrf:       csrfHandler,
		reviews:    reviewsApp,
		search:     searchApp,
		promoCodes: promoCodesGRPC,
	}, nil
}

func (s *Server) setupRoutes() {
	errResolver := errResolveLib.NewErrorStore()

	csrfMiddleware := middlewares.CSRFMiddleware(csrf.NewCscfUsecase(), s.sessions)

	subProd := s.product.InitProductsRoutes()
	subProd.Use(middlewares.AuthMiddleware(s.sessions, errResolver))
	subProd.Use(middlewares.RequestIDMiddleware)
	subProd.Use(csrfMiddleware)
	s.category.InitCategoriesRoutes()

	subCart := s.cart.InitCartRoutes()
	subCart.Use(middlewares.AuthMiddleware(s.sessions, errResolver))
	subCart.Use(middlewares.RequestIDMiddleware)
	subCart.Use(csrfMiddleware)

	subOrder := s.order.InitOrderApp()
	subOrder.Use(middlewares.RequestIDMiddleware)
	subOrder.Use(middlewares.AuthMiddleware(s.sessions, errResolver))
	subOrder.Use(csrfMiddleware)

	s.r.HandleFunc("/api/v1/login", s.auth.LoginUser).Methods(http.MethodPost)
	s.r.HandleFunc("/api/v1/logout", s.sessions.Delete).Methods(http.MethodPost)
	s.r.HandleFunc("/api/v1/signup", s.auth.CreateUser).Methods(http.MethodPost)

	authSub := s.r.Methods(http.MethodGet, http.MethodPost, http.MethodPut).Subrouter()
	authSub.HandleFunc("/api/v1/csrf", s.csrf.GetCsrf).Methods(http.MethodGet)
	authSub.HandleFunc("/api/v1/", s.auth.GetUserById).Methods(http.MethodGet)
	authSub.Use(middlewares.RequestIDMiddleware)
	authSub.Use(middlewares.AuthMiddleware(s.sessions, errResolver))

	s.r.HandleFunc("/api/v1/files/{name}", s.files.GetImage).Methods(http.MethodGet)

	csrfProtected := authSub.Methods(http.MethodGet, http.MethodPost, http.MethodPut).Subrouter()

	csrfProtected.HandleFunc("/api/v1/account", s.profile.GetProfile).Methods(http.MethodGet)
	csrfProtected.HandleFunc("/api/v1/account", s.profile.UpdateProfileData).Methods(http.MethodPut)
	csrfProtected.HandleFunc("/api/v1/change_password", s.profile.ChangePassword).Methods(http.MethodPut)
	csrfProtected.HandleFunc("/api/v1/account/avatar", s.profile.UpdateProfileAvatar).Methods(http.MethodPut)
	csrfProtected.HandleFunc("/api/v1/address", s.address.GetAddress).Methods(http.MethodGet)
	csrfProtected.HandleFunc("/api/v1/address", s.address.UpdateAddressData).Methods(http.MethodPut)
	csrfProtected.HandleFunc("/api/v1/promocodes", s.promoCodes.GetUserPromoCodes).Methods(http.MethodGet)
	csrfProtected.Use(csrfMiddleware)
	csrfProtected.Use(middlewares.RequestIDMiddleware)

	subSearch := s.search.InitSearchRoutes()
	subSearch.Use(middlewares.RequestIDMiddleware)

	reviewsSub := s.reviews.InitRoutes()
	reviewsSub.Use(middlewares.RequestIDMiddleware)
	reviewsSub.Use(middlewares.AuthMiddleware(s.sessions, errResolver))
}

func (s *Server) Run() error {
	s.setupRoutes()

	handler := middlewares.CorsMiddleware(s.r)

	s.log.Info("starting  server", slog.String("address:", s.cfg.Port))
	return http.ListenAndServe(fmt.Sprintf(":%s", s.cfg.Port), handler)
}
