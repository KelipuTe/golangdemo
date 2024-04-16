package client

import (
	"demo-golang/tcp-service/config"
	"fmt"
	"log"
	"testing"
)

func Test_Client_HTTP(t *testing.T) {
	log.Println("version: ", config.Version)

	p1client := NewTCPClient(config.HTTPStr, "127.0.0.1", 9501)
	p1client.SetName(fmt.Sprintf("%s-client", config.HTTPStr))
	p1client.SetDebugStatusOn()
	p1client.Start()
}

func Test_Client_Stream(t *testing.T) {
	log.Println("version: ", config.Version)

	p1client := NewTCPClient(config.StreamStr, "127.0.0.1", 9501)
	p1client.SetName(fmt.Sprintf("%s-client", config.StreamStr))
	p1client.SetDebugStatusOn()
	p1client.Start()
}

func Test_Client_WebSocket(t *testing.T) {
	log.Println("version: ", config.Version)

	p1client := NewTCPClient(config.WebSocketStr, "127.0.0.1", 9501)
	p1client.SetName(fmt.Sprintf("%s-client", config.WebSocketStr))
	p1client.SetDebugStatusOn()
	p1client.Start()
}
