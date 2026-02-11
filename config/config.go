package config

import (
	"os"
	
	"github.com/spf13/viper"
)

type Config struct {
	Host    string `mapstructure:"HOST"`
	Port    string `mapstructure:"PORT"`
	Address string `mapstructure:"ADDRESS"`
	APIKey  string `mapstructure:"API_KEY"`
	
	DBUser 	string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName 	string `mapstructure:"DB_NAME"`
	DBHost 	string `mapstructure:"DB_HOST"`
	DBPort 	string `mapstructure:"DB_PORT"`
}

func LoadConfig() (Config, error) {
	var config Config
	v := viper.New() 

	v.AutomaticEnv()

	keys := []string{
        "HOST",
        "PORT",
		"ADDRESS",
		"API_KEY",
		
        "DB_USER",
        "DB_PASSWORD",
        "DB_NAME",
        "DB_HOST",
        "DB_PORT",
    }

    for _, k := range keys {
        _ = v.BindEnv(k)
    }

	if os.Getenv("APP_ENV") != "production" {
        if _, err := os.Stat(".env"); err == nil {
            v.SetConfigFile(".env")
            _ = v.ReadInConfig()
        }
    }

	err := v.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}