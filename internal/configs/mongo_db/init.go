package mongo_db

import (
	"context"
	"github.com/caarlos0/env"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type mongoConfig struct {
	Username string `env:"MONGO_INITDB_ROOT_USERNAME"`
	Password string `env:"MONGO_INITDB_ROOT_PASSWORD"`
	DBName   string `env:"MONGO_INITDB_DATABASE"`
}

const mongoDb = "mongo_db"

func Connect() (*mongo.Database, error) {
	cfg := mongoConfig{}
	v, err := configs.SetupViper()
	if err != nil {
		return nil, err
	}

	confMap := v.GetStringMap(mongoDb)

	_, err = configs.ParseServiceViperConfig(confMap)
	if err != nil {
		return nil, err
	}

	if err = env.Parse(&cfg); err != nil {
		return nil, err
	}
	//url := fmt.Sprintf("mongodb://%s:%s",
	//	//cfg.Username,
	//	//cfg.Password,
	//	conf.Address,
	//	conf.Port,
	//)

	url := "mongodb://mongo_db:27017"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	log.Printf("connect to mongodb at %s", url)

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}

	db := client.Database(cfg.DBName)

	log.Printf("connected to mongodb at db %s", cfg.DBName)

	return db, nil
}