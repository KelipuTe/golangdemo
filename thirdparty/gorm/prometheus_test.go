package gorm

import (
	gormexprometheus "demo-golang/thirdparty/gormex/callback/prometheus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormprometheus "gorm.io/plugin/prometheus"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

type GormexPrometheusModel struct {
	Id       int64 `gorm:"column:id;primaryKey,autoIncrement"`
	CreateAt int64 `gorm:"column:create_at"`
}

func (t GormexPrometheusModel) TableName() string {
	return "gormex_prometheus"
}

func TestGormexPrometheus(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//DBName==数据库名。用于区分不同的数据库实例
	//RefreshInterval==刷新间隔。客户端刷新指标到内存中的间隔
	//MetricsCollector==指标收集器‌
	//Threads_running==MySQL活跃线程数
	err = db.Use(gormprometheus.New(gormprometheus.Config{
		DBName:          "golangdemo",
		RefreshInterval: 15,
		MetricsCollector: []gormprometheus.MetricsCollector{
			&gormprometheus.MySQL{VariableNames: []string{"Threads_running"}},
		},
	}))
	if err != nil {
		panic(err)
	}

	prometheusCallback := gormexprometheus.NewCallback(prometheus.SummaryOpts{
		Help:      "通过callback监控gorm",
		Namespace: "golangdemo",
		Subsystem: "gorm",
		Name:      "callback",
		ConstLabels: map[string]string{
			"instance_id": "",
		},
		Objectives: map[float64]float64{
			0.5:   0.01,
			0.8:   0.01,
			0.9:   0.01,
			0.99:  0.001,
			0.999: 0.0001,
		},
	})

	err = db.Use(prometheusCallback)
	if err != nil {
		panic(err)
	}

	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		num := rand.Intn(500)
		time.Sleep(time.Duration(num) * time.Millisecond)

		m := GormexPrometheusModel{}
		m.CreateAt = time.Now().UnixMilli()
		err = db.WithContext(ctx).Create(&m).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "TestGormexPrometheus")
			return
		}
		ctx.JSON(http.StatusOK, "TestGormexPrometheus")
	})

	go func() {
		err2 := server.Run(":8080")
		if err2 != nil {
			panic(err2)
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
