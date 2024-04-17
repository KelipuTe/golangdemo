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

	p1innerClient := client.NewTCPClient("127.0.0.1", 9501, config.StreamStr)
	p1innerClient.SetName(fmt.Sprintf("%s-client-user", config.StreamStr))
	p1innerClient.OpenDebug()

	p1innerClient.AfterConnConnect = func(p1conn *client.TCPConnection) {
		if p1innerClient.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.AfterConnConnect", p1innerClient.GetName()))
		}
		P1UserService.RegisteServiceProvider()
	}

	p1innerClient.OnConnGetRequest = func(p1conn *client.TCPConnection) {
		if p1innerClient.IsDebug() {
			fmt.Println(fmt.Sprintf("%s.AfterConnConnect", p1innerClient.GetName()))
		}
		P1UserService.DispatchRequest(p1conn)
	}

	P1UserService.SetInnerClient(p1innerClient)
	p1innerClient.Start()

	signal.WaitForShutdown()
}
