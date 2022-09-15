package main

import (
	"log"
	"net/http"
)

func main() {
	// ServeMux 实现了 Handle 接口
	p7sm := http.NewServeMux()
	// 第二个参数进去会被强转成 HandlerFunc
	// HandlerFunc 实现了 Handle 接口
	p7sm.HandleFunc("/v2", HandleFuncV2)
	p7sm.HandleFunc("/v4", HandleFuncV4)

	p7hs := &http.Server{
		Addr:    "127.0.0.1:9510",
		Handler: p7sm,
	}

	log.Println("http.Server.ListenAndServe...")
	p7hs.ListenAndServe()
}

func HandleFuncV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HandleFuncV2"))
}

func HandleFuncV4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HandleFuncV4"))
}
