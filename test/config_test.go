package test

import (
	"fmt"
	"go-live-chat/config"
	"syscall"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
	config.LoadYumlConfig("../config/app")
	httpPort := viper.GetInt("app.http.port")
	if httpPort == 0 {
		t.Error("load config exception")
	}
}

func TestIncreaseULimit(t *testing.T) {
	config.IncreaseUlimit()
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		t.Error(err)
	}
	fmt.Printf("config %d %d", rLimit.Cur, rLimit.Max)
	if rLimit.Cur != rLimit.Max {
		t.Error("increase ulimit failed")
	}
}
