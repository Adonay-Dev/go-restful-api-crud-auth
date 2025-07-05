package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

func LoadConfiguration() (Config, error) {
	viper.SetConfigFile("../.env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}
