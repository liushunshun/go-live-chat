package main

import (
	"fmt"
	"go-live-chat/config"
	"go-live-chat/internal/helper"
	"go-live-chat/internal/server"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	config.LoadYumlConfig("config/app")
	config.IncreaseUlimit()

	log.Info(fmt.Sprintf("Go Live Chat server start at ip %s config properties: %s", helper.GetServerIp(), helper.ToJsonString(viper.Get("app"))))

	go server.StartWebSocketService()
	go server.StartPprofService()
	server.StartWebService()
}
