package worker

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Task struct {
	C      context.Context
	Input  func(output chan any)
	Output chan any
}

type ClientOptions func(*WorkerClient) error

type WorkerClient struct {
	workers   int
	taskqueue chan *Task

	log    *slog.Logger
	tracer trace.Tracer
}

func WithLogger(l *slog.Logger) ClientOptions {
	return func(c *WorkerClient) error {
		c.log = l
		return nil
	}
}

func WithNumWorkers(num int) ClientOptions {
	return func(c *WorkerClient) error {
		c.workers = num
		return nil
	}
}

func New(opts ...ClientOptions) (*WorkerClient, error) {
	c := WorkerClient{
		log:       slog.Default(),
		workers:   5,
		taskqueue: make(chan *Task),
		tracer:    otel.Tracer("Worker"),
	}

	for _, f := range opts {
		if err := f(&c); err != nil {
			return nil, err
		}
	}

	for i := 0; i <= c.workers; i++ {
		go c.worker(i, c.taskqueue)
	}

	return &c, nil
}

func (w *WorkerClient) Add(in *Task) {
	go func() {
		w.taskqueue <- in
	}()
}

func (w WorkerClient) worker(num int, task chan *Task) {
	for {
		t := <-task
		_, span := w.tracer.Start(t.C, "Worker-"+fmt.Sprint(num))
		if w.log != nil {
			w.log.Info("Running task", slog.Int("worker", num))
		}
		t.Input(t.Output)
		span.AddEvent("Completed task")

		span.End()
	}
}
