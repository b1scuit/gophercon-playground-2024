package tracker

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type TrackerOptions func(*Tracker)

type Tracker struct {
	tp *sdktrace.TracerProvider
}

func New(opts ...TrackerOptions) *Tracker {
	t := Tracker{}

	for _, f := range opts {
		f(&t)
	}

	return &t
}

func (t *Tracker) Setup(ctx context.Context) error {
	exp, err := newExporter(ctx)
	if err != nil {
		return fmt.Errorf("otel setup fail : %w", err)
	}

	// Create a new tracer provider with a batch span processor and the given exporter.
	t.tp = newTraceProvider(exp)

	// Handle shutdown properly so nothing leaks.
	otel.SetTracerProvider(t.tp)
	return nil
}

func (t *Tracker) Shutdown(ctx context.Context) error {
	return t.tp.Shutdown(ctx)
}

func newExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	// Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
	return otlptracehttp.New(ctx)
}
func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("Service"),
		),
	)

	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}
