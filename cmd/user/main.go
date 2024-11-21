package main

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/logger"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	userRepoLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/repository/user"
	sessionsServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/sessions"
	userServiceLib "github.com/go-park-mail-ru/2024_2_kotyari/internal/usecase/user"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"github.com/joho/godotenv"
	"log"
)

const configFile = ".env"

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

	sessionRepo := session
	sessionService := sessionsServiceLib.SessionService{}
	userRepo := userRepoLib.NewUsersStore(dbPool, slogLog)
	userService := userServiceLib.NewUserService(userRepo, inputValidator, slogLog)

	delivery := profile.NewProfilesGrpc(profileRepo, profileService, slogLog)

	app := profilesGrpc.NewProfilesApp(slogLog, delivery)

	err = app.Run("0.0.0.0:8003")
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
