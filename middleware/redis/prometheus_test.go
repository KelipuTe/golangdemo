package redis

import (
	"context"
	redisprometheushook "demo-golang/middleware/redis/hook/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"
)

var (
	addr      = "localhost:6379"
	keyPrefix = "golangdemo:middleware:redis"
)

func TestPrometheus(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	labels := []string{"cmd", "pattern"}
	vector := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Help:      "通过hook监控redis",
			Namespace: "golangdemo",
			Subsystem: "middleware",
			Name:      "redis",
			ConstLabels: map[string]string{
				"instance_id": "",
			},
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

	rdb.AddHook(redisprometheushook.NewHook(vector))

	ctx := context.Background()
	result, err := rdb.FlushDB(ctx).Result()
	if err != nil {
		panic(err)
	}
	t.Log("FlushDB", result, err)

	for i := 0; i < 50; i++ {
		ctx2 := context.Background()
		num2 := rand.Intn(100)
		key2 := keyPrefix + ":" + strconv.Itoa(num2)
		result2, err2 := rdb.Set(ctx2, key2, num2, 5*time.Minute).Result()
		t.Log("Set", key2, num2, result2, err2)
	}

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			ctx2 := context.Background()

			//附加业务信息，让hook里面可以进行区分
			ctx2 = context.WithValue(ctx, "pattern", keyPrefix+":%d")

			num2 := rand.Intn(100)
			key2 := keyPrefix + ":" + string(rune(num2))
			result2, err2 := rdb.Get(ctx2, key2).Result()
			t.Log("Get", key2, num2, result2, err2)
		}
	}()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err2 := http.ListenAndServe(":8081", nil)
		if err2 != nil {
			panic(err2)
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
