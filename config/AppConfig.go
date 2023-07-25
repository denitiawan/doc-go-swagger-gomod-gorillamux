package config

import (
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {
	// APP
	ServerPort string `mapstructure:"SERVER_PORT"`

	// DB
	DBHost       string `mapstructure:"DB_HOST"`
	DBUsername   string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBDriverName string `mapstructure:"DB_DRIVER_NAME"`

	// JWT
	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}

func LoadConfig(path string, env string) (config AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	if env == "dev" {
		viper.SetConfigName("app_dev")
	} else {
		viper.SetConfigName("app_docker")
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
