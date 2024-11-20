package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	errResolveLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/rating_updater/app"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/product"
	"log"
)

func main() {
	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		log.Fatal(err)
	}

	initLogger := logger.InitLogger()
	productsRepo := product.NewProductsStore(dbPool, initLogger)
	errResolver := errResolveLib.NewErrorStore()

	fmt.Println("AAAAAAAAAAAAAAAAA")

	grpcApp := app.NewApp(productsRepo, initLogger, errResolver)
	err = grpcApp.Run()
	if err != nil {
		log.Fatal(err)
	}
}
