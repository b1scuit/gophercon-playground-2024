package proto

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Corer interface {
	DoThething(context.Context) error
}

type ClientOptions func(*ProtoClient) error

func WithCore(c Corer) ClientOptions {
	return func(hc *ProtoClient) error {
		hc.core = c
		return nil
	}
}

type ProtoClient struct {
	UnsafeExampleServerServer

	core   Corer
	tracer trace.Tracer
}

func New(opts ...ClientOptions) (*ProtoClient, error) {
	c := ProtoClient{
		tracer: otel.Tracer("gRPC Server"),
	}

	for _, f := range opts {
		if err := f(&c); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

func (p *ProtoClient) GetExample(ctx context.Context, in *Example) (*Example, error) {
	ctx, span := p.tracer.Start(ctx, "Example")
	defer span.End()

	err := p.core.DoThething(ctx)

	return &Example{
		Example: "Hello world",
	}, err
}
