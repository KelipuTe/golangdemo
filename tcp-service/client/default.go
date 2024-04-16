package client

import "fmt"

const defaultName string = "default-client" //默认名称

func defaultOnClientStart(p1service *TCPClient) {
	if p1service.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.OnServiceStart", p1service.name))
	}
}

func defaultOnClientError(p1service *TCPClient, err error) {
	if p1service.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.OnServiceError", p1service.name))
	}
	fmt.Println(fmt.Sprintf("%s", err))
}

func defaultOnConnConnect(p1conn *TCPConnection) {
	if p1conn.p1client.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.OnConnConnect", p1conn.p1client.name))
	}
}

func defaultOnConnRequest(p1conn *TCPConnection) {
	if p1conn.p1client.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.OnConnRequest", p1conn.p1client.name))
	}
}

func defaultOnConnClose(p1conn *TCPConnection) {
	if p1conn.p1client.IsDebug() {
		fmt.Println(fmt.Sprintf("%s.OnConnClose", p1conn.p1client.name))
	}
}
