package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the service configuration
type Config struct {
	Server Server `mapstructure:"server"`
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
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
		fmt.Println("Using default configuration as config file not found")
	} else {
		fmt.Println("Successfully loaded config file:", viper.ConfigFileUsed())
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("cfg: %+v\n", cfg)

	return &cfg
}
