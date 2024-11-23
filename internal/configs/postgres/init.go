package postgres

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	MainDBCFG = "main_db"
	CSATDBCFG = "csat_db"

	defaultMaxConnections            = 90
	defaultMinConnections            = 0
	defaultMaxConnectionLifeTime     = time.Hour * 2
	defaultMinConnectionIdleLifeTime = time.Minute * 30
)

type postgresVarsConfig struct {
	UsernameVar  string `mapstructure:"db_username"`
	PasswordVar  string `mapstructure:"db_password"`
	DBNameVar    string `mapstructure:"db_name"`
	DockerDBName string `mapstructure:"docker_name"`
}

type postgresConfig struct {
	Username     string
	Password     string
	DBName       string
	DockerDbName string
}

func loadPGConfig(configName string) (postgresConfig, error) {
	v, err := configs.SetupViper()
	if err != nil {
		log.Println("Error loading viper", err.Error())

		return postgresConfig{}, err
	}

	if err = godotenv.Load(configs.EnvPath); err != nil {
		log.Println("Error loading .env", err.Error())
		return postgresConfig{}, err
	}

	var cfgVars postgresVarsConfig
	if err = v.Sub(configName).Unmarshal(&cfgVars); err != nil {
		log.Println("Viper unmarshalling error", err.Error())

		return postgresConfig{}, err
	}

	cfg, err := parseEnvVars(cfgVars)
	if err != nil {
		log.Println("Error loading .env vars", err.Error())

		return postgresConfig{}, err
	}
	log.Printf("postgres config load success")

	return cfg, nil
}

func parseEnvVars(varsConfig postgresVarsConfig) (postgresConfig, error) {
	dbUsername := os.Getenv(varsConfig.UsernameVar)
	dbPassword := os.Getenv(varsConfig.PasswordVar)
	dbName := os.Getenv(varsConfig.DBNameVar)
	dockerDBName := os.Getenv(varsConfig.DockerDBName)

	return postgresConfig{
		Username:     dbUsername,
		Password:     dbPassword,
		DBName:       dbName,
		DockerDbName: dockerDBName,
	}, nil
}

func newPostgresConfigURL(p postgresConfig) string {
	link := "postgres://%s:%s@%s/%s"
	//link := "postgres://%s:%s@localhost:54320/%s"
	return fmt.Sprintf(link,
		url.QueryEscape(p.Username),
		url.QueryEscape(p.Password),
		p.DockerDbName,
		p.DBName,
	)
}

func LoadPgxPool(configName string) (*pgxpool.Pool, error) {
	cfg, err := loadPGConfig(configName)
	if err != nil {
		return nil, err
	}

	poolCfg, err := pgxpool.ParseConfig(newPostgresConfigURL(cfg))
	if err != nil {
		return nil, fmt.Errorf("[LoadPgxPool] failed to parse config: %w", err)
	}

	poolCfg.MaxConns = defaultMaxConnections
	poolCfg.MinConns = defaultMinConnections
	poolCfg.MaxConnLifetime = defaultMaxConnectionLifeTime
	poolCfg.MaxConnIdleTime = defaultMinConnectionIdleLifeTime

	connPool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, fmt.Errorf("[LoadPgxPool] failed to connect to postgres: %w", err)
	}

	//err = testPing(connPool)
	//if err != nil {
	//	return nil, err
	//}

	return connPool, nil
}

//func testPing(pool *pgxpool.Pool) error {
//	if err := pool.Ping(context.Background()); err != nil {
//		return fmt.Errorf("[TestPing] failed to ping postgres: %w", err)
//	}
//
//	return nil
//}
