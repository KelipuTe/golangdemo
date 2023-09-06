package http

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type HTTPServiceV2 struct{}

func (*HTTPServiceV2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTPServiceV2"))
}

type HTTPServiceV4 struct{}

func (*HTTPServiceV4) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTPServiceV4"))
}

func HandleFuncV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HandleFuncV2"))
}

func HandleFuncV4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HandleFuncV4"))
}

func TestMode02(p7s6t *testing.T) {
	// ServeMux 实现了 http.Handle 接口
	p7sm := http.NewServeMux()
	// 第二个参数是 http.Handle 接口的实例
	p7sm.Handle("/v2", &HTTPServiceV2{})
	p7sm.Handle("/v4", &HTTPServiceV4{})

	p7hs := &http.Server{
		Addr:    "127.0.0.1:9511",
		Handler: p7sm,
	}

	log.Println("http.Server.ListenAndServe...")
	err := p7hs.ListenAndServe()
	fmt.Println(err)
}

func TestMode04(p7s6t *testing.T) {
	p7sm := http.NewServeMux()
	// 第二个参数会被强转成 http.HandlerFunc 类型
	// http.HandlerFunc 类型实现了 http.Handle 接口
	p7sm.HandleFunc("/v2", HandleFuncV2)
	p7sm.HandleFunc("/v4", HandleFuncV4)

	p7hs := &http.Server{
		Addr:    "127.0.0.1:9511",
		Handler: p7sm,
	}

	log.Println("http.Server.ListenAndServe...")
	err := p7hs.ListenAndServe()
	fmt.Println(err)
}

func TestMode06(p7s6t *testing.T) {
	log.Println("http.ListenAndServe...")
	// 第二个参数是 http.Handle 接口的实例
	err := http.ListenAndServe("127.0.0.1:9511", &HTTPServiceV2{})
	fmt.Println(err)
}

func TestMode08(p7s6t *testing.T) {
	// 直接使用 http 包里的 DefaultServeMux
	http.HandleFunc("/v2", HandleFuncV2)
	http.HandleFunc("/v4", HandleFuncV4)

	log.Println("http.ListenAndServe...")
	// 第二个参数传 nil
	err := http.ListenAndServe("127.0.0.1:9511", nil)
	fmt.Println(err)
}
