package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

// Builder gin普罗米修斯监控中间件
// 监控，请求总数，正在处理的请求数，请求时间
// \Namespace+Subsystem+Name，是唯一的即可，要不然会panic
// 比如，可以是，公司部门-部门管理的服务-从服务采集的什么数据
type Builder struct {
	// 这客户端是干什么的
	Help string
	// 命名空间
	Namespace string
	// 子系统
	Subsystem string
	// 名字
	Name string
	// 客户端实例id
	InstanceId string
}

func (t *Builder) TotalRequestCount() gin.HandlerFunc {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Help:      t.Help,
		Namespace: t.Namespace,
		Subsystem: t.Subsystem,
		Name:      t.Name + "_total_req",
		ConstLabels: map[string]string{
			"instance_id": t.InstanceId,
		},
	})
	prometheus.MustRegister(counter)

	return func(ctx *gin.Context) {
		counter.Inc()

		ctx.Next()
	}
}

func (t *Builder) ActiveRequestGauge() gin.HandlerFunc {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Help:      t.Help,
		Namespace: t.Namespace,
		Subsystem: t.Subsystem,
		Name:      t.Name + "_active_req",
		ConstLabels: map[string]string{
			"instance_id": t.InstanceId,
		},
	})
	prometheus.MustRegister(gauge)

	return func(ctx *gin.Context) {
		gauge.Inc()
		defer gauge.Dec()

		ctx.Next()
	}
}

func (t *Builder) ResponseTimeSummary() gin.HandlerFunc {
	labels := []string{"method", "pattern", "status"}
	vector := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Help:      t.Help,
			Namespace: t.Namespace,
			Subsystem: t.Subsystem,
			Name:      t.Name + "_resp_time",
			ConstLabels: map[string]string{
				"instance_id": t.InstanceId,
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

	return func(ctx *gin.Context) {
		startTime := time.Now()
		defer func() {
			timeSince := time.Since(startTime).Milliseconds()
			method := ctx.Request.Method
			pattern := ctx.FullPath()
			status := ctx.Writer.Status()
			vector.WithLabelValues(method, pattern, strconv.Itoa(status)).
				Observe(float64(timeSince))
		}()

		ctx.Next()
	}
}
