package server

import (
	"fmt"
	"go-live-chat/common/helper"
	"go-live-chat/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StartWebServer() {
	router := gin.Default()
	config.InitWebRouters(router)

	if err := http.ListenAndServe(fmt.Sprint(helper.GetServerIp(), ":", viper.GetInt("http.port")), router); err != nil {
		log.Fatalf("web server failed: %v", err)
	}
}
