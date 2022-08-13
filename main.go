package main

import (
	"go-live-chat/config"
	"go-live-chat/server/gateway/websocket"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.Info("gochat server starting...")

	//初始化配置数据
	config.Init()
	log.Info("config init success")

	//websocket监听端口
	websocketPort := viper.GetInt("websocket.port")
	websocket.Start(websocketPort)

	log.Info("gochat server started at: 127.0.0.1:", websocketPort)
}
