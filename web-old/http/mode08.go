package http

import (
	"fmt"
	"log"
	"net/http"
)

func Mode08() {
	log.Println("http.ListenAndServe...")
	// 第二个参数直接就是 Handle 接口的实例
	err := http.ListenAndServe("127.0.0.1:9510", &HTTPServiceV2{})
	fmt.Println(err)
}
