package postgres

import (
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

const (
	defaultMaxConnections            = 90
	defaultMinConnections            = 0
	defaultMaxConnectionLifeTime     = time.Hour * 2
	defaultMinConnectionIdleLifeTime = time.Minute * 30
)

type postgresConfig struct {
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
}

func mustLoadPGConfig() postgresConfig {
	cfg := postgresConfig{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	test := postgresConfig{}
	if test == cfg {
		log.Fatal("[mustLoadPGConfig] postgres config is empty")
	}

	log.Printf("postgres config load success")

	return cfg
}

func newPostgresConfigURL(p postgresConfig) string {
	return fmt.Sprintf("postgres://%s:%s@pg_db/%s",
		p.Username,
		p.Password,
		p.DBName,
	)
}

func MustLoadPgxPool() *pgxpool.Pool {
	cfg := mustLoadPGConfig()

	poolCfg, err := pgxpool.ParseConfig(newPostgresConfigURL(cfg))
	if err != nil {
		log.Fatalf("[MustLoadPgxPool] failed to parse config: %s", err.Error())
	}

	poolCfg.MaxConns = defaultMaxConnections
	poolCfg.MinConns = defaultMinConnections
	poolCfg.MaxConnLifetime = defaultMaxConnectionLifeTime
	poolCfg.MaxConnIdleTime = defaultMinConnectionIdleLifeTime

	connPool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		log.Fatalf("[MustLoadPgxPool] failed to connect to postgres: %s", err.Error())
	}

	testPing(connPool)

	return connPool
}

func testPing(pool *pgxpool.Pool) {
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("[TestPing] failed to ping postgres: %s", err.Error())
	}
}
