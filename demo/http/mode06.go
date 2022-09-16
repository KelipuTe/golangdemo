package http

import (
	"log"
	"net/http"
)

func Mode06() {
	// 如果不显式的创建 ServeMux，http 包里有个默认的 DefaultServeMux
	http.HandleFunc("/v2", HandleFuncV2)
	http.HandleFunc("/v4", HandleFuncV4)

	log.Println("http.ListenAndServe...")
	// 第二个参数传 nil 的时候，里面就会使用默认的那个 ServeMux
	http.ListenAndServe("127.0.0.1:9510", nil)
}
