package zipkin

import (
	"context"
	optlex "demo-golang/middleware/opentelemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"testing"
	"time"
)

func TestZipkin(t *testing.T) {
	//初始化opentelemetry
	res := optlex.NewResource()
	exporter := optlex.NewZipkinExporter()
	traceProvider := optlex.NewTraceProvider(res, exporter)
	defer traceProvider.Shutdown(context.Background())
	otel.SetTracerProvider(traceProvider)
	prop := optlex.NewPropagation()
	otel.SetTextMapPropagator(prop)

	//塞一些数据给zipkin
	{
		ctx := context.Background()

		//这个Tracer就是上面的traceProvider生成的
		//name参数，用来标记trace是在哪里生成的，方便找
		tracer := otel.Tracer("middlewareoptlzipkin")

		//最顶层（比如，某一次请求）
		ctx1, span1 := tracer.Start(ctx, "span1") //从ctx派生
		defer span1.End()
		span1.SetAttributes(attribute.String("span1", "value1"))
		time.Sleep(time.Millisecond * 200) //模拟业务耗时

		//trace id这样拿出来
		t.Log("trace id", span1.SpanContext().TraceID())

		//最顶层的子事件（比如，调用下游服务）。这里写了两个子事件。
		{
			//最顶层的第一个子事件
			span1.AddEvent("span11")
			ctx11, span11 := tracer.Start(ctx1, "span11") //从最顶层ctx1派生
			defer span11.End()
			span11.SetAttributes(attribute.String("span11", "value11"))
			time.Sleep(time.Millisecond * 200) //模拟业务耗时
			//第一个子事件的子事件（比如，下游服务调用数据库）
			{
				span11.AddEvent("span111")
				_, span111 := tracer.Start(ctx11, "span111") //从第一个子事件ctx11派生
				defer span111.End()
				span111.SetAttributes(attribute.String("span111", "value111"))
				time.Sleep(time.Millisecond * 200) //模拟业务耗时
			}

			//最顶层的第二个子事件
			span1.AddEvent("span12")
			_, span12 := tracer.Start(ctx1, "span12") //从最顶层ctx1派生
			defer span12.End()
			span12.SetAttributes(attribute.String("span12", "value12"))
			time.Sleep(time.Millisecond * 200) //模拟业务耗时
		}
	}
}
