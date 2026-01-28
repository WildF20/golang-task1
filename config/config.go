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