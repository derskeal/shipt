package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Envs EnvType

type EnvType struct {
	Port string

	DbHost     string
	DbPort     int
	DbName     string
	DbUsername string
	DbPassword string
}

func init() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Config file error: %w \n", err))
	}

	Envs = EnvType{
		Port:       viper.GetString("PORT"),
		DbHost:     viper.GetString("DB_HOST"),
		DbPort:     viper.GetInt("DB_PORT"),
		DbName:     viper.GetString("DB_NAME"),
		DbUsername: viper.GetString("DB_USERNAME"),
		DbPassword: viper.GetString("DB_PASSWORD"),
	}
}
