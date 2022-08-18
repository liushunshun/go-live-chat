package server

import (
	"fmt"
	"go-live-chat/internal/helper"
	"net/http"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/lesismal/llib/std/crypto/tls"
	"github.com/lesismal/nbio/nbhttp"
	"github.com/lesismal/nbio/nbhttp/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	qps   uint64 = 0
	total uint64 = 0

	server *nbhttp.Server
)

func newUpgrader() *websocket.Upgrader {
	u := websocket.NewUpgrader()
	u.OnMessage(func(c *websocket.Conn, messageType websocket.MessageType, data []byte) {
		c.WriteMessage(messageType, data)
		atomic.AddUint64(&qps, 1)
	})
	return u
}

func onWebsocket(w http.ResponseWriter, r *http.Request) {
	upgrader := newUpgrader()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	wsConn := conn.(*websocket.Conn)
	wsConn.SetReadDeadline(time.Time{})
}

func StartWebSocketService() {
	rsaCertPEM := helper.ReadBytes("../../config/cer.pem")
	rsaKeyPEM := helper.ReadBytes("../../config/privatekey.pem")
	cert, err := tls.X509KeyPair(rsaCertPEM, rsaKeyPEM)
	if err != nil {
		log.Fatalf("tls.X509KeyPair failed: %v", err)
	}
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()

	mux := &http.ServeMux{}
	mux.HandleFunc("/wss", onWebsocket)

	server = nbhttp.NewServer(nbhttp.Config{
		Network:                 "tcp",
		AddrsTLS:                GetWsAddrFromConfig(),
		TLSConfig:               tlsConfig,
		MaxLoad:                 1000000,
		ReleaseWebsocketPayload: true,
		Handler:                 mux,
	})

	err = server.Start()
	if err != nil {
		fmt.Printf("nbio.Start failed: %v\n", err)
		return
	}
	defer server.Stop()

	ticker := time.NewTicker(time.Second * 10)
	for i := 1; true; i++ {
		<-ticker.C
		n := atomic.SwapUint64(&qps, 0)
		total += n
		fmt.Printf("running for %v seconds, NumGoroutine: %v, qps: %v, total: %v\n", i*10, runtime.NumGoroutine(), n, total)
	}
}

func GetWsAddrFromConfig() []string {
	portConfig := viper.GetString("app.websocket.port")
	ip := helper.GetServerIp()
	var addrs []string
	if !strings.Contains(portConfig, "-") {
		addr := ip + ":" + portConfig
		addrs = append(addrs, addr)
	} else {
		tagIndex := strings.Index(portConfig, "-")
		startPort := helper.ToInt(portConfig[0:tagIndex])
		endPort := helper.ToInt(portConfig[tagIndex+1:])

		for i := startPort; i < endPort; i++ {
			addr := fmt.Sprint(ip, ":", i)
			addrs = append(addrs, addr)
		}
	}
	return addrs
}
