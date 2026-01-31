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

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}