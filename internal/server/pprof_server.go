package server

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"go-live-chat/internal/helper"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartPprofService() {
	if err := http.ListenAndServe(fmt.Sprint(helper.GetServerIp(), ":", viper.GetInt("app.pprof.port")), nil); err != nil {
		log.Fatalf("Pprof failed: %v", err)
	}
}
