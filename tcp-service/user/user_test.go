package user

import (
	"demo-golang/tcp-service/client"
	"demo-golang/tcp-service/config"
	"demo-golang/tcp-service/tool/signal"
	"fmt"
	"log"
	"testing"
)

func Test_User(t *testing.T) {
	log.Println("version: ", config.Version)

	P1UserService := &UserService{}

	p1innerClient := client.NewTCPClient(config.StreamStr, "127.0.0.1", 9501)
	p1innerClient.SetName(fmt.Sprintf("%s-client-user", config.StreamStr))
	p1innerClient.OpenDebug()

	p1innerClient.OnClientStart = func(p1client *client.TCPClient) {
		if p1client.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnServiceStart", p1client.GetName()))
		}
		P1UserService.SetInnerClient(p1client)
	}

	p1innerClient.OnConnConnect = func(p1conn *client.TCPConnection) {
		if p1innerClient.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnConnConnect", p1innerClient.GetName()))
		}
		P1UserService.RegisteServiceProvider()
	}

	p1innerClient.OnConnRequest = func(p1conn *client.TCPConnection) {
		if p1innerClient.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.OnConnConnect", p1innerClient.GetName()))
		}
		P1UserService.DispatchRequest(p1conn)
	}
	p1innerClient.Start()

	signal.WaitForShutdown()
}
