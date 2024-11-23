package configs

import "github.com/spf13/viper"

const (
	KeyAddress      = "address"
	KeyPort         = "port"
	ServicesConfigs = "services"
	ConfigPath      = "configs"
	EnvPath         = ".env"
)

func SetupViper() (*viper.Viper, error) {
	viper.AddConfigPath(ConfigPath)
	viper.SetConfigName(ServicesConfigs)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return viper.GetViper(), nil
}
