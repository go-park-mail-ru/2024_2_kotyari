package main

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/app/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	user2 "github.com/go-park-mail-ru/2024_2_kotyari/internal/grpc_api/user"
	userRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	userServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/joho/godotenv"
	"log"
)

const configFile = ".env"

// todo вынос в app
func main() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPool, err := postgres.LoadPgxPool()
	if err != nil {
		log.Fatalf("не инициализируется бд %v", err)
	}

	slogLog := logger.InitLogger()

	// todo добавить
	inputValidator := utils.NewInputValidator()

	userRepo := userRepoLib.NewUsersStore(dbPool, slogLog)
	userService := userServiceLib.NewUserService(userRepo, inputValidator, slogLog)

	delivery := user2.NewUsersGrpc(userService, userRepo, slogLog)

	app := user.NewUsersApp(slogLog, delivery)

	err = app.Run("0.0.0.0:8001")
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
