package opentelemetry

import (
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"time"
)

func NewResource() *resource.Resource {
	// resource，用于区分不同服务

	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("golangdemo"),
			semconv.ServiceVersion("v0.0.1"),
		),
	)
	if err != nil {
		panic(err)
	}

	return res
}

func NewZipkinExporter() *zipkin.Exporter {
	// exporter，会将trace数据发送到监控平台
	// 这里创建的是zipkin的exporter，会将trace数据发送到zipkin

	exporter, err := zipkin.New("http://localhost:9411/api/v2/spans")
	if err != nil {
		panic(err)
	}
	return exporter
}

func NewTraceProvider(res *resource.Resource, exporter *zipkin.Exporter) *trace.TracerProvider {
	// 初始化trace provider，用来在打点的时候构建trace
	// trace.WithBatcher==批量发送
	// trace.WithBatchTimeout==批量发送的超时时间，默认5s
	// trace.WithResource==附加资源信息

	return trace.NewTracerProvider(
		trace.WithResource(res),
		trace.WithBatcher(
			exporter,
			trace.WithBatchTimeout(3*time.Second),
		),
	)
}

func NewPropagation() propagation.TextMapPropagator {
	// 配置跨服务调用时的携带哪些上下文数据。
	// propagation.TraceContext==传递trace数据（如，trace id）
	// propagation.Baggage==传递自定义数据（如，用户id）

	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
