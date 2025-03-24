package gorm

import (
	gormexprometheus "demo-golang/thirdparty/gormex/middleware/prometheus"
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
	dsn := "root:root@tcp(localhost:3306)/golangdemo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.Use(gormprometheus.New(gormprometheus.Config{
		//数据库名。用于区分不同的数据库实例
		DBName: "golangdemo",
		//刷新间隔。客户端刷新指标到内存中的间隔
		RefreshInterval: 15,
		//指标收集器‌
		MetricsCollector: []gormprometheus.MetricsCollector{
			//Threads_running==MySQL活跃线程数
			&gormprometheus.MySQL{VariableNames: []string{"Threads_running"}},
		},
	}))
	if err != nil {
		panic(err)
	}

	prometheusCallback := gormexprometheus.NewCallback(prometheus.SummaryOpts{
		Help:      "监控gorm的callback接口",
		Namespace: "golangdemo",
		Subsystem: "gormprometheus",
		Name:      "gormcallback",
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
		var num int

		minute := time.Now().Second()
		if minute > 40 {
			num = rand.Intn(500)
		} else if minute > 20 {
			num = rand.Intn(300)
		} else {
			num = rand.Intn(100)
		}

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
