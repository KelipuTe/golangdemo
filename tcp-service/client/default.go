package client

import "fmt"

const defaultName string = "default-client" //默认名称

func defaultOnClientError(client *TCPClient, err error) {
	if client.IsDebug() {
		fmt.Println(fmt.Sprintf("client [%s] OnClientError", client.name))
	}
	fmt.Println(fmt.Sprintf("client [%s] err:[%s]", client.name, err.Error()))
}

func defaultAfterConnConnect(conn *TCPConnection) {
	if conn.belongToClient.IsDebug() {
		fmt.Println(fmt.Sprintf("client [%s] AfterConnConnect", conn.belongToClient.name))
	}
}

func defaultOnConnGetRequest(conn *TCPConnection) {
	if conn.belongToClient.IsDebug() {
		fmt.Println(fmt.Sprintf("client [%s] OnConnGetRequest", conn.belongToClient.name))
	}
}

func defaultAfterConnClose(conn *TCPConnection) {
	if conn.belongToClient.IsDebug() {
		fmt.Println(fmt.Sprintf("client [%s] AfterConnClose", conn.belongToClient.name))
	}
}
