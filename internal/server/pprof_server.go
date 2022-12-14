package server

import (
	"fmt"
	"go-live-chat/internal/helper"
	"net/http"
	_ "net/http/pprof"

	log "github.com/lesismal/nbio/logging"
	"github.com/spf13/viper"
)

func StartPprofService() {
	if err := http.ListenAndServe(fmt.Sprint(helper.GetServerIp(), ":", viper.GetInt("app.pprof.port")), nil); err != nil {
		log.Error("Pprof failed: %v", err)
	}
}
