package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("http.ListenAndServe...")
	// 第二个参数直接就是 Handle 接口的实例
	http.ListenAndServe("127.0.0.1:9510", &HTTPService2{})
}

type HTTPService2 struct{}

func (*HTTPService2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTPService"))
}
