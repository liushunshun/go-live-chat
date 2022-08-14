package server

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"go-live-chat/common/helper"

	"github.com/spf13/viper"
)

func StartPprofServer() {
	if err := http.ListenAndServe(fmt.Sprint(helper.GetServerIp(), ":", viper.GetInt("pprof.port")), nil); err != nil {
		log.Fatalf("Pprof failed: %v", err)
	}
}
