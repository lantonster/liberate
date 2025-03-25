package config

import (
	"context"

	"github.com/lantonster/liberate/pkg/color"
	"github.com/lantonster/liberate/pkg/database"
	"github.com/lantonster/liberate/pkg/log"
	"github.com/spf13/viper"
)

// Config holds the service configuration
type Config struct {
	Server Server         `mapstructure:"server"`
	MySQL  database.MySQL `mapstructure:"mysql"`
	Logger log.Config     `mapstructure:"logger"`
}

type Server struct {
	Port int    `mapstructure:"port" default:"8080"`
	Host string `mapstructure:"host" default:"0.0.0.0"`
}

// LoadConfig loads the configuration from environment variables and config file
func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	log.SetLogger(cfg.Logger)
	log.WithContext(context.Background()).Info(color.Green.Sprint("config loaded"))

	return &cfg
}
