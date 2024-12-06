package configs

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/spf13/viper"
	"log/slog"
)

const (
	KeyDomain       = "domain"
	KeyAddress      = "address"
	KeyPort         = "port"
	ServicesConfigs = "services"
	ConfigPath      = "configs"
	EnvPath         = ".env"
)

type ServiceViperConfig struct {
	Domain  string
	Address string
	Port    string
}

type KafkaConfig struct {
	Domain string
	Port   string
}

func SetupViper() (*viper.Viper, error) {
	viper.AddConfigPath(ConfigPath)
	viper.SetConfigName(ServicesConfigs)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return viper.GetViper(), nil
}

func ParseServiceViperConfig(config map[string]any) (ServiceViperConfig, error) {
	domain, ok := config[KeyDomain].(string)
	if !ok {
		slog.Error("[ParseServiceViperConfig] Failed to parse domain")

		return ServiceViperConfig{}, errs.FailedToParseConfig
	}

	address, ok := config[KeyAddress].(string)
	if !ok {
		slog.Error("[ParseServiceViperConfig] Failed to parse address")

		return ServiceViperConfig{}, errs.FailedToParseConfig
	}

	port, ok := config[KeyPort].(string)
	if !ok {
		slog.Error("[ParseServiceViperConfig] Failed to parse port")

		return ServiceViperConfig{}, errs.FailedToParseConfig
	}

	return ServiceViperConfig{
		Domain:  domain,
		Address: address,
		Port:    port,
	}, nil
}

func ParseKafkaViperConfig(config map[string]any) (KafkaConfig, error) {
	domain, ok := config[KeyDomain].(string)
	if !ok {
		slog.Error("[ParseServiceViperConfig] Failed to parse domain")

		return KafkaConfig{}, errs.FailedToParseConfig
	}

	port, ok := config[KeyPort].(string)
	if !ok {
		slog.Error("[ParseServiceViperConfig] Failed to parse port")

		return KafkaConfig{}, errs.FailedToParseConfig
	}

	return KafkaConfig{
		Domain: domain,
		Port:   port,
	}, nil
}
