package main

import (
	"context"
	"cpuprofile/core"
	"cpuprofile/proto"
	"cpuprofile/tracker"
	"cpuprofile/worker"
	"log/slog"
	"net"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	now := time.Now()
	slog.Info("Loading client")
	t := tracker.New()
	t.Setup(ctx)

	defer func() { t.Shutdown(ctx) }()

	workers, _ := worker.New()

	client, _ := core.New(
		core.WithWorker(workers),
	)
	protoCore, _ := proto.New(
		proto.WithCore(client),
	)

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		slog.Error("Error starting server", slog.Any("error", err))
		os.Exit(1)
	}

	server := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	proto.RegisterExampleServerServer(server, protoCore)

	slog.Info("Started server", slog.Duration("time", time.Since(now)))

	if err := server.Serve(listener); err != nil {
		slog.Error("Error running server", slog.Any("error", err))
		os.Exit(1)
	}
}
