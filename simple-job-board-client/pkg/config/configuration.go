package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	// HTTPServer
	HTTPServerHost string `mapstructure:"http_server_host"`
	HTTPServerPort string `mapstructure:"http_server_port"`

	// Postgres
	PostgresHost     string `mapstructure:"postgres_host"`
	PostgresPort     string `mapstructure:"postgres_port"`
	PostgresUser     string `mapstructure:"postgres_user"`
	PostgresPassword string `mapstructure:"postgres_password"`
	PostgresDatabase string `mapstructure:"postgres_database"`
	PostgresSSLMode  string `mapstructure:"postgres_ssl_mode"`
}

func NewConfiguration(filename string) (*Configuration, error) {
	viper.SetConfigType("yaml")

	if mode := viper.GetString("MODE"); mode == "PROD" {
		filename = filename + ".prod.yml"
	} else {
		filename = filename + ".dev.yml"
	}
	viper.SetConfigName(filename)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Load configuration from file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal configuration into struct
	var config Configuration
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
