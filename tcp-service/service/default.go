package service

import (
	"fmt"
)

const defaultName string = "default-service" //tcp服务端默认名称

func defaultOnServiceStart(service *TCPService) {
	if service.IsDebug() {
		fmt.Println(fmt.Sprintf("service [%s] OnServiceStart", service.name))
	}
}

func defaultOnServiceError(service *TCPService, err error) {
	if service.IsDebug() {
		fmt.Println(fmt.Sprintf("service [%s] OnServiceError", service.name))
	}
	fmt.Println(err.Error())
}

func defaultOnConnConnect(conn *TCPConnection) {
	if conn.belongToService.IsDebug() {
		fmt.Println(fmt.Sprintf("service [%s] OnConnConnect", conn.belongToService.name))
	}
}

func defaultOnConnRequest(conn *TCPConnection) {
	if conn.belongToService.IsDebug() {
		fmt.Println(fmt.Sprintf("service [%s] OnConnRequest", conn.belongToService.name))
	}
}

func defaultOnConnClose(conn *TCPConnection) {
	if conn.belongToService.IsDebug() {
		fmt.Println(fmt.Sprintf("service [%s] OnConnClose", conn.belongToService.name))
	}
}
