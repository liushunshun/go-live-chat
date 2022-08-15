// gobwas/ws
// 1. Zero-copy upgrade
// 2. No intermediate allocations during I/O
// 3. Low-level API which allows to build your own logic of packet handling and buffers reuse
// 4. High-level wrappers and helpers around API in wsutil package, which allow to start fast without digging the protocol internals

package server

import (
	"fmt"
	"go-live-chat/internal/helper"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartWebSocketService() {
	port := viper.GetInt("app.websocket.port")
	http.HandleFunc(viper.GetString("app.websocket.url"), handler)
	address := fmt.Sprint(helper.GetServerIp(), ":", port)
	http.ListenAndServe(address, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 升级协议
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		return
	}
	log.Info("新链接请求，建立连接:", conn.RemoteAddr().String())
	reader(conn)
}

func reader(conn net.Conn) {
	if msg, op, err := wsutil.ReadClientData(conn); err != nil {
		conn.Close()
	} else {
		// This is commented out since in demo usage, stdout is showing messages sent from > 1M connections at very high rate
		fmt.Println("收到消息：" + string(msg) + " 来自：" + conn.RemoteAddr().String())

		response := []byte("Hi client")

		err = wsutil.WriteServerMessage(conn, op, response)
		fmt.Println("服务端回消息：" + string(response) + " 给：" + conn.RemoteAddr().String())

		if err != nil {
			log.Println(err)
			return
		}
	}
}
