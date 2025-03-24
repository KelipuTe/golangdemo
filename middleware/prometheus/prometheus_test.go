package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"testing"
	"time"
)

func TestPrometheus(t *testing.T) {
	//业务服务本体
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("TestPrometheus"))
		})

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()

	//另外起一个端口给服务端拉数据用
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			panic(err)
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
