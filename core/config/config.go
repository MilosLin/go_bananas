package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

var _instance *viper.Viper

type Config struct {
}

func Instance() *viper.Viper {
	if _instance == nil {
		_instance = new()
	}
	return _instance
}

func new() *viper.Viper {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		err = errors.New("No configuration file loaded")
		log.Fatal(err)
	}
	return viper.GetViper()
}
