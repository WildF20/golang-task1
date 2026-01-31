package config

import (
	"os"
	"strings"
	
	"github.com/spf13/viper"
)

type Config struct {
	Host    string `mapstructure:"HOST"`
	Port    string `mapstructure:"PORT"`
	Address string `mapstructure:"ADDRESS"`
	
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
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

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