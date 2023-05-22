package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBConnection            string        `mapstructure:"DB_CONNECTION"`
	ServerPort              string        `mapstructure:"SERVER_PORT"`
	LogLevel                string        `mapstructure:"LOG_LEVEL"`
	AccessTokenKey          string        `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey         string        `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration    time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	CloudinaryCloudName     string        `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudinaryApiKey        string        `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinaryApiSecret     string        `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudinaryUploadFolder  string        `mapstructure:"CLOUDINARY_UPLOAD_FOLDER"`
	PaginateDefaultPage     int           `mapstructure:"PAGINATE_DEFAULT_PAGE"`
	PaginateDefaultPageSize int           `mapstructure:"PAGINATE_DEFAULT_PAGE_SIZE"`
	PaginateDefaultType     string        `mapstructure:"PAGINATE_DEFAULT_TYPE"`
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
