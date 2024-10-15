package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	Server ServerConfig
}

func NewConfig() (Config, error) {
	viper.AddConfigPath("./services/gateway/cmd")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	return Config{

		Server: ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetInt("server.port"),
		},
	}, nil
}
