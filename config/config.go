package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
}

var configPath = "/etc/app/"

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
