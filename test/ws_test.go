package test

import (
	"fmt"
	"go-live-chat/config"
	"go-live-chat/internal/server"
	"testing"
)

func TestLoadWsPort(t *testing.T) {
	config.LoadYumlConfig("../config/app")
	wsports := server.GetWsAddrFromConfig()
	for _, element := range wsports {
		fmt.Println(string(element))
	}
	if len(wsports) != 50 {
		t.Error("load ws port config exception")
	}
}
