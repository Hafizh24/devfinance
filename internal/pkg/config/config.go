package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBConnection         string        `mapstructure:"DB_CONNECTION"`
	ServerPort           string        `mapstructure:"SERVER_PORT"`
	LogLevel             string        `mapstructure:"LOG_LEVEL"`
	AccessTokenKey       string        `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey      string        `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(fileconfigpath string) (Config, error) {
	var config Config

	viper.AddConfigPath(fileconfigpath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	// nolint:errcheck
	viper.Unmarshal(&config)
	return config, nil
}
