package gin

import (
	"demo-golang/thirdparty/ginex/middleware/prometheus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestGinPrometheus(t *testing.T) {
	server := gin.Default()
	prometheusBuilder := &prometheus.Builder{
		Help:      "监控gin的http接口",
		Namespace: "golangdemo",
		Subsystem: "gin",
		Name:      "http",
	}

	server.Use(
		prometheusBuilder.TotalRequestCount(),
		prometheusBuilder.ActiveRequestGauge(),
		prometheusBuilder.ResponseTimeSummary(),
	)

	server.GET("/", func(ctx *gin.Context) {
		num := rand.Intn(500)
		time.Sleep(time.Duration(num) * time.Millisecond)

		ctx.JSON(http.StatusOK, "TestGinPrometheus")
	})

	go func() {
		err := server.Run(":8080")
		if err != nil {
			panic(err)
		}
	}()

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
