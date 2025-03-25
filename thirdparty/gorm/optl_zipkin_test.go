package gorm

import (
	"context"
	optlex "demo-golang/middleware/opentelemetry"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestOptlZipkin(t *testing.T) {
	//初始化opentelemetry
	res := optlex.NewResource()
	exporter := optlex.NewZipkinExporter()
	traceProvider := optlex.NewTraceProvider(res, exporter)
	defer traceProvider.Shutdown(context.Background())
	otel.SetTracerProvider(traceProvider)
	prop := optlex.NewPropagation()
	otel.SetTextMapPropagator(prop)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 配置gorm的opentelemetry插件
	// tracing.WithoutMetrics==不要记录metrics
	// tracing.WithoutQueryVariables==不要记录查询参数
	err = db.Use(tracing.NewPlugin(
		tracing.WithQueryFormatter(func(query string) string {
			t.Log(query) //语句打出来看看
			return query
		}),
		tracing.WithDBName("golangdemo"),
		tracing.WithoutMetrics(),
		tracing.WithoutQueryVariables(),
	))
	if err != nil {
		panic(err)
	}

	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		tracer := otel.Tracer("gormoptlzipkin")

		ctx1, span1 := tracer.Start(ctx.Request.Context(), "gin")
		defer span1.End()
		span1.SetAttributes(attribute.String("gin", "gin"))

		num := rand.Intn(500)
		time.Sleep(time.Duration(num) * time.Millisecond)

		m := GormexPrometheusModel{}
		m.CreateAt = time.Now().UnixMilli()
		err = db.WithContext(ctx1).Create(&m).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "TestOptlZipkin")
			return
		}
		ctx.JSON(http.StatusOK, "TestOptlZipkin")
	})

	go func() {
		err2 := server.Run(":8080")
		if err2 != nil {
			panic(err2)
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
