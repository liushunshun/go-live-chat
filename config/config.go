package config

import (
	"syscall"

	"github.com/spf13/viper"
)

func LoadYumlConfig(path string) {
	viper.SetConfigName(path)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func IncreaseUlimit() {
	// Increase resources limitations
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
}
