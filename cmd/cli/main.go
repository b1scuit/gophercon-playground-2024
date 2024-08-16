package main

import (
	"context"
	"cpuprofile/core"
	"cpuprofile/worker"
	"flag"
	"log/slog"
	"os"
	"time"
)

func main() {
	var verbose bool
	var things int
	flag.BoolVar(&verbose, "verbose", true, "Should slog print verbosly")
	flag.IntVar(&things, "things", 100, "Added things to do")
	flag.Parse()

	now := time.Now()

	var l *slog.Logger
	if verbose {
		l = slog.Default()
	}

	workers, _ := worker.New(
		worker.WithLogger(l),
	)
	client, _ := core.New(
		core.WithWorker(workers),
		core.WithThingsToDo(things),
	)

	slog.Info("Running application")

	if err := client.DoThething(context.Background()); err != nil {
		slog.Error("Error running the code")
		os.Exit(1)
	}

	slog.Info("Application successfully run", slog.Duration("time", time.Since(now)))
}
