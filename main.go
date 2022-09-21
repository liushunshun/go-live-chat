package main

import (
	"flag"
	"fmt"
	"go-live-chat/config"
	"go-live-chat/internal/helper"
	"go-live-chat/internal/server"

	log "github.com/lesismal/nbio/logging"
	"github.com/spf13/viper"
)

var (
	debug = flag.Bool("pprof", false, "start pprof server")
)

func main() {
	config.LoadYumlConfig("config/app")
	config.IncreaseUlimit()

	log.Info(fmt.Sprintf("Go Live Chat server start at ip %s config properties: %s", helper.GetServerIp(), helper.ToJsonString(viper.Get("app"))))

	if *debug {
		go server.StartPprofService()
	}
	go server.StartWebSocketService()
	server.StartWebService()
}
