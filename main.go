package main

import (
	"go-live-chat/config"
	"go-live-chat/server"
)

func main() {
	config.Init()

	go server.StartWebSocketServer()
	go server.StartPprofServer()
	server.StartWebServer()
}
