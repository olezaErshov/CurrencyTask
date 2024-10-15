package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Host string
	Port int
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

func NewConfig() (Config, error) {
	viper.AddConfigPath("./services/currency/cmd")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	return Config{
		DB: DBConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			Username: viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
			DBName:   viper.GetString("postgres.dbname"),
			SSLMode:  viper.GetString("postgres.sslmode"),
		},

		Server: ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetInt("server.port"),
		},
	}, nil
}
