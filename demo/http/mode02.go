package http

import (
	"log"
	"net/http"
)

func Mode02() {
	// ServeMux 实现了 Handle 接口
	p7sm := http.NewServeMux()
	// 第二个参数直接就是 Handle 接口的实例
	p7sm.Handle("/v2", &HTTPServiceV2{})
	p7sm.Handle("/v4", &HTTPServiceV4{})

	p7hs := &http.Server{
		Addr:    "127.0.0.1:9510",
		Handler: p7sm,
	}

	log.Println("http.Server.ListenAndServe...")
	p7hs.ListenAndServe()
}

type HTTPServiceV2 struct{}

func (*HTTPServiceV2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTPServiceV2"))
}

type HTTPServiceV4 struct{}

func (*HTTPServiceV4) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTPServiceV4"))
}
