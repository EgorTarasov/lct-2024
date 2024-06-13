package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

// NewConsoleExporter заглушка для проверки работы трейсера.
func NewConsoleExporter() (oteltrace.SpanExporter, error) {
	return stdouttrace.New()
}

// NewOTLPExporter создает экспортер трейсов в OTLP.
func NewOTLPExporter(ctx context.Context, endpoint string) (oteltrace.SpanExporter, error) {
	insecureOpt := otlptracehttp.WithInsecure()

	endpointOpt := otlptracehttp.WithEndpoint(endpoint)

	return otlptracehttp.New(ctx, insecureOpt, endpointOpt)
}

// NewTraceProvider фабрика для создания провайдера трейсера.
func NewTraceProvider(exp oteltrace.SpanExporter, appName string) *oteltrace.TracerProvider {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(appName),
		),
	)

	if err != nil {
		panic(err)
	}

	return oteltrace.NewTracerProvider(
		oteltrace.WithBatcher(exp),
		oteltrace.WithResource(r),
	)
}
