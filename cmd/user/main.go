package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/app/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	configs "github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	grpc2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/metrics/grpc"
	metrics2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/middlewares/metrics"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net/http"
	"time"
)

const (
	userService = "user_go"
	configFile  = ".env"
)

// todo вынос в app
func main() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	viper, err := configs.SetupViper()
	if err != nil {
		log.Fatalf("Error setting up viper %v", err)
	}

	conf := viper.GetStringMap(userService)

	slogLog := logger.InitLogger()

	errorResolver := errs.NewErrorStore()

	metrics, err := grpc2.NewGrpcMetrics("user")
	if err != nil {
		slogLog.Error("Ошибка при регистрации метрики", slog.String("error", err.Error()))
	}

	interceptor := metrics2.NewGrpcMiddleware(*metrics, errorResolver)
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.ServerMetricsInterceptor))

	app, err := user.NewUsersApp(slogLog, server, conf)

	router := mux.NewRouter()
	router.PathPrefix("/metrics").Handler(promhttp.Handler())
	serverProm := http.Server{Handler: router, Addr: fmt.Sprintf(":%d", 8081), ReadHeaderTimeout: 10 * time.Second}

	go func() {
		if err := serverProm.ListenAndServe(); err != nil {
			log.Println("fail auth.ListenAndServe")
		}
	}()

	err = app.Run()
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
