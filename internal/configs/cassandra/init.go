package cassandra

import (
	"errors"
	"github.com/gocql/gocql"
	"log"

	"github.com/caarlos0/env"
)

const (
	cassandraHost         = "cassandra"
	cassandraKeyspaceName = "oxic"
)

type cassandraConfig struct {
	Username string `env:"CASSANDRA_USER"`
	Password string `env:"CASSANDRA_PASSWORD"`
}

func loadCassandraConfig() (cassandraConfig, error) {
	cfg := cassandraConfig{}

	if err := env.Parse(&cfg); err != nil {
		return cassandraConfig{}, err
	}

	emptyConfig := cassandraConfig{}
	if cfg == emptyConfig {
		return cassandraConfig{}, errors.New("[loadCassandraConfig] cassandra config is empty")
	}

	log.Printf("cassandra config load success")

	return cfg, nil
}

func LoadCassandraCluster() (*gocql.Session, error) {
	cfg, err := loadCassandraConfig()
	if err != nil {
		return nil, err
	}

	cluster := gocql.NewCluster(cassandraHost)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.Username,
		Password: cfg.Password,
	}
	cluster.Keyspace = cassandraKeyspaceName
	cluster.Consistency = gocql.LocalQuorum
	cassandraSession, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return cassandraSession, nil
}
