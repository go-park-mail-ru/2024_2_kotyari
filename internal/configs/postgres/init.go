package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v5/pgxpool"
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

func loadPGConfig() (postgresConfig, error) {
	cfg := postgresConfig{}

	if err := env.Parse(&cfg); err != nil {
		return postgresConfig{}, err
	}

	emptyConfig := postgresConfig{}
	if cfg == emptyConfig {
		return postgresConfig{}, errors.New("[loadPGConfig] postgres config is empty")
	}

	log.Printf("postgres config load success")

	return cfg, nil
}

func newPostgresConfigURL(p postgresConfig) string {
	link := "postgres://%s:%s@pg_db/%s"
	return fmt.Sprintf(link,
		url.QueryEscape(p.Username),
		url.QueryEscape(p.Password),
		p.DBName,
	)
}

func LoadPgxPool() (*pgxpool.Pool, error) {
	cfg, err := loadPGConfig()
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

	err = testPing(connPool)
	if err != nil {
		return nil, err
	}

	return connPool, nil
}

func testPing(pool *pgxpool.Pool) error {
	if err := pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("[TestPing] failed to ping postgres: %w", err)
	}

	return nil
}
