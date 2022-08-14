package server

import (
	"fmt"
	"go-live-chat/common/helper"
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartWebSocketServer() {
	port := viper.GetInt("websocket.port")
	http.HandleFunc(viper.GetString("websocket.url"), handler)
	address := fmt.Sprint(helper.GetServerIp(), ":", port)
	http.ListenAndServe(address, nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		log.Info("新链接请求，升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)
		return
	}
	log.Info("新链接请求，建立连接:", conn.RemoteAddr().String())

	reader(conn)
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("收到消息：" + string(p) + " 来自：" + conn.RemoteAddr().String())

		response := []byte("Hi client")

		err = conn.WriteMessage(websocket.TextMessage, response)

		fmt.Println("服务端回消息：" + string(response) + " 给：" + conn.RemoteAddr().String())

		if err != nil {
			log.Println(err)
			return
		}

	}
}
