package main

import (
	"context"
	"cpuprofile/core"
	"cpuprofile/handler"
	"cpuprofile/tracker"
	"cpuprofile/worker"
	"log/slog"
)

func main() {
	ctx := context.Background()
	t := tracker.New()
	t.Setup(ctx)

	defer func() { t.Shutdown(ctx) }()

	// Finally, set the tracer that can be used for this package.

	workers, _ := worker.New()

	client, _ := core.New(
		core.WithWorker(workers),
	)

	httpHandler, _ := handler.New(
		handler.WithCore(client),
		handler.WithRoutes(),
	)

	if err := httpHandler.Do(); err != nil {
		slog.Error("Failed to start HTTP server", slog.Any("error", err))
	}
}
