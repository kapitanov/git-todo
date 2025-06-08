package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kapitanov/git-todo/internal/commands"
)

var (
	version = "develop"
)

func main() {
	cmd := commands.New()
	cmd.Version = version

	err := withSigterm(func(ctx context.Context) error {
		return cmd.ExecuteContext(ctx)
	})
	if err != nil {
		cmd.PrintErrf("%s\n", err.Error())
		os.Exit(-1)
	}
}

func withSigterm(fn func(ctx context.Context) error) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-signals
		cancel()
	}()

	return fn(ctx)
}
