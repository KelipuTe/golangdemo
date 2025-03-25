package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var (
	help      = "prometheus测试"
	namespace = "golangdemo"
	subsystem = "middleware"
	name      = "prometheus"
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

	//塞一些数据给prometheus
	go func() {

		//Counter
		counter := prometheus.NewCounter(prometheus.CounterOpts{
			Help:      help,
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name + "_counter",
			ConstLabels: map[string]string{
				"instance_id": "",
			},
		})
		prometheus.MustRegister(counter)

		//Gauge
		gauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Help:      help,
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name + "_gauge",
			ConstLabels: map[string]string{
				"instance_id": "",
			},
		})
		prometheus.MustRegister(gauge)

		gauge.Inc()
		gauge.Dec()

		//Summary
		labels := []string{"label1", "label2"}
		vector := prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Help:      help,
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      name + "_summary",
				ConstLabels: map[string]string{
					"instance_id": "",
				},
				//key是百分比，val是误差。
				//key=0.5；val=0.01。那0.49-0.51都算。
				//误差不能太大，一般是比百分比多一位小数。
				Objectives: map[float64]float64{
					0.5:   0.01,
					0.7:   0.01,
					0.9:   0.01,
					0.99:  0.001,
					0.999: 0.0001,
				},
			},
			labels,
		)
		prometheus.MustRegister(vector)

		for {
			time.Sleep(100 * time.Millisecond)
			t.Log("塞一些数据给prometheus")

			counter.Inc()

			num1 := rand.Intn(10)
			if time.Now().UnixMilli()%2 == 0 {
				gauge.Inc()
				gauge.Add(float64(num1))
			} else {
				gauge.Dec()
				gauge.Sub(float64(num1))
			}

			num2 := rand.Intn(500)
			if time.Now().UnixMilli()%2 == 0 {
				vector.WithLabelValues("label1", "label21").Observe(float64(num2))
			} else {
				vector.WithLabelValues("label1", "label22").Observe(float64(num2))
			}
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
