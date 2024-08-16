package core

import (
	"context"
	"cpuprofile/worker"
	"testing"
)

func BenchmarkDoTheThingNoLog(b *testing.B) {
	workers, _ := worker.New(
		worker.WithLogger(nil),
	)
	coreClient, _ := New(
		WithWorker(workers),
	)
	ctx := context.Background()

	for i := 0; i <= b.N; i++ {
		coreClient.DoThething(ctx)
	}
}
func BenchmarkDoTheThingLogger(b *testing.B) {
	workers, _ := worker.New()
	coreClient, _ := New(
		WithWorker(workers),
	)
	ctx := context.Background()

	for i := 0; i <= b.N; i++ {
		coreClient.DoThething(ctx)
	}
}
