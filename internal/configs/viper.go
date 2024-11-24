package configs

import (
	"github.com/spf13/viper"
)

const (
	KeyAddress      = "address"
	KeyPort         = "port"
	ServicesConfigs = "services"
	ConfigPath      = "configs"
	EnvPath         = ".env"
)

type ServiceViperConfig struct {
	Address string
	Port    string
}

func SetupViper() (*viper.Viper, error) {
	viper.AddConfigPath(ConfigPath)
	viper.SetConfigName(ServicesConfigs)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return viper.GetViper(), nil
}

func ParseServiceViperConfig(config map[string]any) ServiceViperConfig {
	return ServiceViperConfig{
		Address: config[KeyAddress].(string),
		Port:    config[KeyPort].(string),
	}
}
