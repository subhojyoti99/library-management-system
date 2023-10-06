package config

import (
	"github.com/spf13/viper"
)

// AppConfig stores the application configuration.
var AppConfig struct {
	Port        int    // Server Port
	DatabaseURL string // PostgreSQL Database URL
}

// InitConfig initializes the application configuration.
func InitConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	AppConfig.Port = viper.GetInt("PORT")
	AppConfig.DatabaseURL = viper.GetString("DATABASE_URL")
}
