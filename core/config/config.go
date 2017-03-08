/**
 * 設定檔操作類別，固定讀取執行檔路徑下的 api-config.yml
 *
 *  Usage:
 * 	c := config.NewConfig()
 *	ip := c.Setting.GetString("ip")
 */
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
