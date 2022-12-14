package server

import (
	"fmt"
	"go-live-chat/config"
	"go-live-chat/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/lesismal/nbio/logging"
	"github.com/spf13/viper"
)

func StartWebService() {
	router := gin.Default()
	config.InitWebRouters(router)

	if err := http.ListenAndServe(fmt.Sprint(helper.GetServerIp(), ":", viper.GetInt("app.http.port")), router); err != nil {
		log.Error("web server failed: %v", err)
	}
}
