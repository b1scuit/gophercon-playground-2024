package handler

import (
	"context"
	"log/slog"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Corer interface {
	DoThething(context.Context) error
}

type ClientOptions func(*HandlerClient) error

func WithCore(c Corer) ClientOptions {
	return func(hc *HandlerClient) error {
		hc.core = c
		return nil
	}
}

func WithRoutes() ClientOptions {
	return func(c *HandlerClient) error {
		c.handleFunc("/example", c.ExampleHandler)

		return nil
	}
}

type HandlerClient struct {
	mux  *http.ServeMux
	core Corer

	tracer trace.Tracer
}

func New(opts ...ClientOptions) (*HandlerClient, error) {
	c := HandlerClient{
		mux:    http.NewServeMux(),
		tracer: otel.Tracer("Handler"),
	}

	for _, f := range opts {
		if err := f(&c); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

func (hc *HandlerClient) Do() error {
	return http.ListenAndServe(":8080", otelhttp.NewHandler(hc.mux, "/"))
}

func (hc HandlerClient) ExampleHandler(w http.ResponseWriter, r *http.Request) {
	slog.InfoContext(r.Context(), "Starting Example Handler")
	ctx, span := hc.tracer.Start(r.Context(), "HTTP Entrypoint")
	defer span.End()

	hc.core.DoThething(ctx)

	span.SetStatus(codes.Ok, "All Good")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (hc *HandlerClient) handleFunc(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
	hc.mux.Handle(pattern, handler)
}
