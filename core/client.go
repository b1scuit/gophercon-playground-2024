package core

import (
	"context"
	"cpuprofile/worker"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Workeror interface {
	Add(*worker.Task)
}

type ClientOptions func(*CoreClient) error

func WithWorker(w Workeror) ClientOptions {
	return func(c *CoreClient) error {
		c.wp = w
		return nil
	}
}

func WithThingsToDo(things int) ClientOptions {
	return func(c *CoreClient) error {
		c.thingsToDo = things
		return nil
	}
}

type CoreClient struct {
	wp     Workeror
	tracer trace.Tracer

	thingsToDo int
}

func New(opts ...ClientOptions) (*CoreClient, error) {
	c := CoreClient{
		thingsToDo: 5,
		tracer:     otel.Tracer("Core"),
	}

	for _, f := range opts {
		if err := f(&c); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

func (c CoreClient) DoThething(ctx context.Context) error {
	ctx, span := c.tracer.Start(ctx, "Core.DoTheThing")
	defer span.End()

	span.AddEvent("Called to complete tasks")
	for i := 0; i < c.thingsToDo; i++ {
		ctx, span := c.tracer.Start(ctx, "Task")
		span.SetAttributes(attribute.Int("thing_its_doing", i))
		defer span.End()

		response := make(chan any)

		c.wp.Add(&worker.Task{
			C: ctx,
			Input: func(out chan any) {
				out <- true
			},
			Output: response,
		})
		<-response
	}

	return nil
}
