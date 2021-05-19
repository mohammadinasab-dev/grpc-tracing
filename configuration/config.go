package configuration

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBAddress  string `mapstructure:"DB_ADDRESS"`
	DBName     string `mapstructure:"GRPC_DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	config = Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
