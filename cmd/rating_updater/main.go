package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/rating_updater/app"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/rating_updater/app/usecase"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/product"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/reviews"
	"github.com/joho/godotenv"
)

const configFile = ".env"

func main() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		log.Fatal(err)
	}

	initLogger := logger.InitLogger()
	productsRepo := product.NewProductsStore(dbPool, initLogger)
	reviewsRepo := reviews.NewReviewsStore(dbPool, initLogger)
	productUpdaterManager := usecase.NewRatingUpdateService(productsRepo, reviewsRepo, initLogger)

	grpcApp := app.NewApp(productUpdaterManager, initLogger)
	err = grpcApp.Run()
	if err != nil {
		log.Fatal(err)
	}
}
