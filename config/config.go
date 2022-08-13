package config

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config/app")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
