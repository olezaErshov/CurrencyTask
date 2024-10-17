package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string
	Port int
}

type UrlsConfig struct {
	AuthGenerator   string
	CurrencyService string
}

type Config struct {
	Server ServerConfig
	Urls   UrlsConfig
}

func NewConfig() (Config, error) {
	viper.AddConfigPath("config")
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
		Urls: UrlsConfig{
			AuthGenerator:   viper.GetString("url.auth_generator"),
			CurrencyService: viper.GetString("url.currency_service"),
		},
	}, nil
}
